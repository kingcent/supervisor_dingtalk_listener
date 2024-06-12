package notify

import (
	"log"
	"os"

	// "github.com/irebit/dingtalk_group_robot"
	"github.com/irebit/dingding_bot"
)

func DingTalk(title, text, messageUrl string) {
	// log.Println("os.Args[1]", os.Args[1], title, text, messageUrl)
	// robot := dingtalk_group_robot.New()
	robot := dingding_bot.New().SetAccessToken(os.Args[1])
	msg := dingding_bot.NewMarkDown()
	title = ">" + title
	msg.SetContent(title, text)
	ok, err := robot.Send(msg)
	if !ok || err != nil {
		log.Println("failed to send msg. ", msg, err)
	}
}
