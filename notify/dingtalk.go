package notify

import (
	"os"

	"github.com/irebit/dingtalk_group_robot"
)

func DingTalk(title, desc string) {
	dingtalk_group_robot.New().SetAccessToken(os.Args[1]).Send(dingtalk_group_robot.NewLink().SetContent(title, desc, "", ""))
}
