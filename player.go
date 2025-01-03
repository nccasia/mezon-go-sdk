package mezonsdk

import (
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
	userId    string
	username  string

	// TODO: streaming video (#rapchieuphim)
	// videoTrack *webrtc.TrackLocalStaticRTP
	audioTrack *webrtc.TrackLocalStaticSample
	audiences  map[string]*webrtc.PeerConnection
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
}

func NewAudioPlayer(clanId, channelId, userId, username, token string) (AudioPlayer, error) {
	stnConn, err := stn.NewWSConnection(&configs.Config{
		BasePath:     "stn.mezon.vn",
		Timeout:      15,
		InsecureSkip: true,
		UseSSL:       true,
	}, channelId, username, token)

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

	// save to store
	rtcConnection := &streamingRTCConn{
		stnWs:      stnConn,
		clanId:     clanId,
		channelId:  channelId,
		userId:     userId,
		username:   username,
		audioTrack: audioTrack,
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

	if rtcConn.(*streamingRTCConn).peer == nil {
		return
	}

	if rtcConn.(*streamingRTCConn).peer.ConnectionState() != webrtc.PeerConnectionStateClosed {
		rtcConn.(*streamingRTCConn).peer.Close()
	}

	MapStreamingRtcConn.Delete(channelId)
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
		pc, err := c.createPeerConnection(&offer)

		// If there is no error, send a success message
		c.stnWs.SendMessage(&stn.WsMsg{
			ClientId: event.ClientId,
		})

	case "sd_answer":
		var answerSDP string
		err := json.Unmarshal(event.Value, &answerSDP)
		if err != nil {
			return err
		}

		return c.peer.SetRemoteDescription(webrtc.SessionDescription{
			Type: webrtc.SDPTypeAnswer,
			SDP:  answerSDP,
		})

	case "ice_candidate":

		var i webrtc.ICECandidateInit
		err := json.Unmarshal(event.Value, &i)
		if err != nil {
			return err
		}

		return c.addICECandidate(i)
	}

	return nil
}

func (c *streamingRTCConn) sendOffer() error {
	offer, err := c.peer.CreateOffer(nil)
	if err != nil {
		return err
	}
	if err := c.peer.SetLocalDescription(offer); err != nil {
		return err
	}

	byteJson, _ := json.Marshal(offer)

	// send socket signaling, gzip compress data
	return c.stnWs.SendMessage(&stn.WsMsg{
		Key:       "session_publisher",
		ClanId:    c.clanId,
		ChannelId: c.channelId,
		UserId:    c.userId,
		Username:  c.username,
		Value:     byteJson,
	})
}

func (c *streamingRTCConn) onICECandidate(i *webrtc.ICECandidate, clanId, channelId, userId, username string) error {
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
		UserId:    userId,
		Username:  username,
	})
}

func (c *streamingRTCConn) addICECandidate(i webrtc.ICECandidateInit) error {
	return c.peer.AddICECandidate(i)
}

func (c *streamingRTCConn) Play(filePath string) error {

	// Open a OGG file and start reading using our OGGReader
	file, oggErr := os.Open(filePath)
	if oggErr != nil {
		return oggErr
	}

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
	for ; true; <-ticker.C {
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
	return nil
}

func (c *streamingRTCConn) createPeerConnection(offer *webrtc.SessionDescription) error {
	pc, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return err
	}
	if err := pc.SetRemoteDescription(*offer); err != nil {
		return err
	}
	rtpSender, err := pc.AddTrack(c.audioTrack)
	if err != nil {
		return err
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
		case webrtc.ICEConnectionStateConnected:
			// TODO: event ice connected
			jsonData, _ := json.Marshal(map[string]string{"ChannelId": c.channelId})
			c.stnWs.SendMessage(&stn.WsMsg{
				ClanId:    c.clanId,
				ChannelId: c.channelId,
				Key:       "connect_publisher",
				Value:     jsonData,
				UserId:    c.userId,
			})
		case webrtc.ICEConnectionStateClosed:
			rtcConn, ok := MapStreamingRtcConn.Load(c.channelId)
			if !ok {
				return
			}

			if rtcConn.(*streamingRTCConn).peer.ConnectionState() != webrtc.PeerConnectionStateClosed {
				rtcConn.(*streamingRTCConn).peer.Close()
			}

			MapStreamingRtcConn.Delete(channelId)
		}
	})
	peerConnection.OnICECandidate(func(i *webrtc.ICECandidate) {
		rtcConnection.onICECandidate(i, channelId, clanId, userId, username)
	})

	// send offer
	rtcConnection.sendOffer()
}
