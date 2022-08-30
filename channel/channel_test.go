package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelDemo_LogFromChannel(t *testing.T) {
	//Test Channel Demo
	channelA := ChannelDemo{
		make(chan string),
		"1000",
	}

	channelB := ChannelDemo{
		make(chan string),
		"1001",
	}

	defer func() {
		close(channelA.Ch)
		close(channelB.Ch)
		fmt.Println("close all channel")
	}()

	go func() {
		time.Sleep(3 * time.Second)
		channelA.SendMsgToChannel(channelB.Ch, "i love u!!")
	}()

	channelB.LogFromChannel()
}
