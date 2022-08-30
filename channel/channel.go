package channel

import "fmt"

//demo for learning channel

type ChannelDemo struct {
	Ch     chan string
	DemoID string
}

func (c *ChannelDemo) LogFromChannel() {
	fmt.Println("LogFromChannel！！")
	//receive message
	msg := <-c.Ch
	fmt.Println("My DemoID is", c.DemoID, "receive msg from channel:", msg)
}

func (c *ChannelDemo) SendMsgToChannel(ch chan string, msg string) {
	fmt.Println("SendMsgToChannel！！")
	if ch == nil {
		fmt.Println("channel is nil!")
	}
	//send message
	ch <- msg

}
