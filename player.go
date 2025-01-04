package mezonsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"

	"os"
	"sync"
	"time"

	"github.com/nccasia/mezon-go-sdk/configs"
	"github.com/nccasia/mezon-go-sdk/stn"
	"github.com/pion/webrtc/v4"
	"github.com/pion/webrtc/v4/pkg/media"
	"github.com/pion/webrtc/v4/pkg/media/oggreader"
)

var (
	MapStreamingRtcConn sync.Map // map[channelId]*RTCConnection
)

type streamingRTCConn struct {
	stnWs stn.IWSConnection

	clanId    string
	channelId string

	// TODO: streaming video (#rapchieuphim)
	// videoTrack *webrtc.TrackLocalStaticRTP
	audioTrack *webrtc.TrackLocalStaticSample
	audiences  map[string]*webrtc.PeerConnection

	ctx        context.Context
	cancelFunc context.CancelFunc
}

var config = webrtc.Configuration{
	ICEServers: []webrtc.ICEServer{
		{
			URLs:       []string{"turn:turn.mezon.vn:5349", "stun:stun.l.google.com:19302"},
			Username:   "turnmezon",
			Credential: "QuTs4zUEcbylWemXL7MK",
		},
	},
}

type AudioPlayer interface {
	Play(filePath string) error
	Close(channelId string)
	Cancel(channelId string)
}

func NewAudioPlayer(clanId, channelId, userId, username, token string) (AudioPlayer, error) {
	stnConn, err := stn.NewWSConnection(&configs.Config{
		BasePath:     "localhost:8081", //"stn.mezon.vn",
		Timeout:      15,
		InsecureSkip: true,
		UseSSL:       false,
	}, channelId, username, token)
	if err != nil {
		return nil, err
	}

	// // Create a video track
	// videoTrack, err := webrtc.NewTrackLocalStaticRTP(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeVP8}, fmt.Sprintf("video_vp8_%s", channelId), fmt.Sprintf("video_vp8_%s", channelId))
	// if err != nil {
	// 	return nil, err
	// }
	// _, err = peerConnection.AddTrack(videoTrack)
	// if err != nil {
	// 	return nil, err
	// }

	// Create a audio track
	audioTrack, err := webrtc.NewTrackLocalStaticSample(webrtc.RTPCodecCapability{MimeType: webrtc.MimeTypeOpus}, fmt.Sprintf("audio_opus_%s", channelId), fmt.Sprintf("audio_opus_%s", channelId))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	// save to store
	rtcConnection := &streamingRTCConn{
		stnWs:      stnConn,
		clanId:     clanId,
		channelId:  channelId,
		audioTrack: audioTrack,
		ctx:        ctx,
		cancelFunc: cancel,
	}

	// ws receive message handler ( on event )
	stnConn.SetOnMessage(rtcConnection.OnWebsocketEvent)
	MapStreamingRtcConn.Store(channelId, rtcConnection)

	rtcConnection.audioTrack = audioTrack

	return rtcConnection, nil
}

func (c *streamingRTCConn) Close(channelId string) {
	rtcConn, ok := MapStreamingRtcConn.Load(channelId)
	if !ok {
		return
	}

	for _, peer := range rtcConn.(*streamingRTCConn).audiences {
		if peer != nil && peer.ConnectionState() != webrtc.PeerConnectionStateClosed {
			peer.Close()
		}
	}
	rtcConn.(*streamingRTCConn).cancel()

	MapStreamingRtcConn.Delete(channelId)
}

func (c *streamingRTCConn) Cancel(channelId string) {
	rtcConn, ok := MapStreamingRtcConn.Load(channelId)
	if !ok {
		return
	}
	rtcConn.(*streamingRTCConn).cancel()
}

func (c *streamingRTCConn) cancel() {
	if c.cancelFunc != nil {
		c.cancelFunc()
	}

	c.ctx, c.cancelFunc = context.WithCancel(context.Background())
}

func (c *streamingRTCConn) OnWebsocketEvent(event *stn.WsMsg) error {

	switch event.Key {
	case "session_subscriber":
		// receive offer from subscriber
		var offer webrtc.SessionDescription
		err := json.Unmarshal(event.Value, &offer)
		if err != nil {
			return err
		}
		_, err = c.createPeerConnection(&offer, event.ClientId)
		if err != nil {
			return err
		}

	case "ice_candidate":

		var i webrtc.ICECandidateInit
		err := json.Unmarshal(event.Value, &i)
		if err != nil {
			return err
		}

		return c.addICECandidate(i, event.ClientId)
	}

	return nil
}

func (c *streamingRTCConn) sendAnswer(pc *webrtc.PeerConnection) error {
	answer, err := pc.CreateAnswer(nil)
	if err != nil {
		return err
	}
	if err := pc.SetLocalDescription(answer); err != nil {
		return err
	}

	byteJson, _ := json.Marshal(answer)

	return c.stnWs.SendMessage(&stn.WsMsg{
		Key:       "sd_answer",
		ClanId:    c.clanId,
		ChannelId: c.channelId,
		Value:     byteJson,
	})
}

func (c *streamingRTCConn) onICECandidate(i *webrtc.ICECandidate, clanId, channelId, clientId string) error {
	if i == nil {
		return nil
	}
	// If you are serializing a candidate make sure to use ToJSON
	// Using Marshal will result in errors around `sdpMid`
	candidateString, err := json.Marshal(i.ToJSON())
	if err != nil {
		return err
	}

	return c.stnWs.SendMessage(&stn.WsMsg{
		Key:       "ice_candidate",
		Value:     candidateString,
		ClanId:    clanId,
		ChannelId: channelId,
		ClientId:  clientId,
	})
}

func (c *streamingRTCConn) addICECandidate(i webrtc.ICECandidateInit, clientId string) error {
	var err error
	if peer, ok := c.audiences[clientId]; ok {
		err = peer.AddICECandidate(i)
	}
	return err
}

func (c *streamingRTCConn) Play(filePath string) error {

	// Open a OGG file and start reading using our OGGReader
	file, oggErr := os.Open(filePath)
	if oggErr != nil {
		return oggErr
	}
	defer file.Close()

	// Open on oggfile in non-checksum mode.
	ogg, _, oggErr := oggreader.NewWith(file)
	if oggErr != nil {
		return oggErr
	}

	// Keep track of last granule, the difference is the amount of samples in the buffer
	var lastGranule uint64

	// It is important to use a time.Ticker instead of time.Sleep because
	// * avoids accumulating skew, just calling time.Sleep didn't compensate for the time spent parsing the data
	// * works around latency issues with Sleep (see https://github.com/golang/go/issues/44343)
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()

	c.cancel()

	for {
		select {
		case <-c.ctx.Done():
			return nil
		case <-ticker.C:
			pageData, pageHeader, oggErr := ogg.ParseNextPage()
			if errors.Is(oggErr, io.EOF) {
				log.Println("All audio pages parsed and sent")
				return nil
			}

			if oggErr != nil {
				return oggErr
			}

			// The amount of samples is the difference between the last and current timestamp
			sampleCount := float64(pageHeader.GranulePosition - lastGranule)
			lastGranule = pageHeader.GranulePosition
			sampleDuration := time.Duration((sampleCount/48000)*1000) * time.Millisecond

			if oggErr = c.audioTrack.WriteSample(media.Sample{Data: pageData, Duration: sampleDuration}); oggErr != nil {
				return oggErr
			}
		}
	}
}

func (c *streamingRTCConn) createPeerConnection(offer *webrtc.SessionDescription, clientId string) (*webrtc.PeerConnection, error) {
	pc, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return nil, err
	}
	if err := pc.SetRemoteDescription(*offer); err != nil {
		return nil, err
	}
	rtpSender, err := pc.AddTrack(c.audioTrack)
	if err != nil {
		return nil, err
	}

	// Read incoming RTCP packets
	// Before these packets are returned they are processed by interceptors. For things
	// like NACK this needs to be called.
	go func() {
		rtcpBuf := make([]byte, 1500)
		for {
			if _, _, rtcpErr := rtpSender.Read(rtcpBuf); rtcpErr != nil {
				return
			}
		}
	}()

	pc.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		log.Printf("Connection State has changed %s \n", state.String())

		switch state {
		case webrtc.ICEConnectionStateConnected, webrtc.ICEConnectionStateClosed:
			c.stnWs.SendMessage(&stn.WsMsg{
				ClanId:    c.clanId,
				ChannelId: c.channelId,
				Key:       "session_state_changed",
				ClientId:  clientId,
				State:     int(state),
			})
		}
	})

	pc.OnICECandidate(func(i *webrtc.ICECandidate) {
		c.onICECandidate(i, c.channelId, c.clanId, clientId)
	})

	c.sendAnswer(pc)

	return pc, nil
}
