package notify

import (
	"os"

	"github.com/irebit/dingtalk_group_robot"
)

func DingTalk(title, text, messageUrl string) {
	// log.Println("os.Args[1]", os.Args[1], title, text, messageUrl)
	dingtalk_group_robot.New().SetAccessToken(os.Args[1]).Send(dingtalk_group_robot.NewLink().SetContent(title, text, messageUrl, ""))
}
