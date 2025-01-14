package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mezonsdk "github.com/nccasia/mezon-go-sdk"
	swagger "github.com/nccasia/mezon-go-sdk/mezon-api"
	"github.com/nccasia/mezon-go-sdk/mezon-protobuf/mezon/v2/common/api"
	"github.com/nccasia/mezon-go-sdk/mezon-protobuf/mezon/v2/common/rtapi"

	"github.com/antihax/optional"
)

func main() {
	player, err := mezonsdk.NewAudioPlayer("123456", "123456", "123456", "komu", "token")
	if err != nil {
		fmt.Println("error", err)
		return
	}

	// Cleanup function (called when program stops)
	cleanup := func() {
		fmt.Println("Cleaning up...")
		player.Close("123456")
	}

	// Catch the program stop signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := player.Play("output.ogg"); err != nil {
			fmt.Println("Error playing audio:", err)
			// End if music playback fails
			stop <- syscall.SIGTERM
		}
	}()

	// Wait for stop signal
	<-stop
	fmt.Println("Stopping program...")

	// Call cleanup function before exiting
	cleanup()

	return // stop testing

	client, err := mezonsdk.NewClient("xxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	if err != nil {
		panic(err)
	}

	socket, err := client.CreateSocket()

	if err != nil {
		panic(err)
	}

	socket.SetOnPong(func(e *rtapi.Envelope) error {
		fmt.Printf("pong => cid: %s, message: %+v \n", e.Cid, e.GetPong())
		return nil
	})

	time.Sleep(1 * time.Second)

	socket.SetOnChannelMessage(func(e *rtapi.Envelope) error {
		fmt.Printf("messages => cid: %s, message: %+v \n", e.Cid, e.GetChannelMessage())
		return nil
	})

	err = socket.SendMessage(&rtapi.Envelope{
		Message: &rtapi.Envelope_ChannelMessageSend{
			ChannelMessageSend: &rtapi.ChannelMessageSend{
				ClanId:           "1827955317304987648",
				ChannelId:        "1827955317309181955",
				Mode:             2,
				Content:          "{\"t\":\"Test test test\"}",
				Mentions:         []*api.MessageMention{},
				Attachments:      []*api.MessageAttachment{},
				AnonymousMessage: false,
				MentionEveryone:  false,
				Avatar:           "",
				IsPublic:         true,
				Code:             0,
			},
		},
	})
	if err != nil {
		panic(err)
	}

	data, _, err := client.Api.MezonListChannelVoiceUsers(context.Background(), &swagger.MezonListChannelVoiceUsersOpts{
		ClanId:      optional.NewString("1827955317304987648"),
		ChannelType: optional.NewInt32(4),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("data: %+v \n", data)

	select {}
}
