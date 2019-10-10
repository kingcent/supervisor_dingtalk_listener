package notify

import (
	"log"
	"os"

	"github.com/irebit/dingtalk_group_robot"
)

func DingTalk(title, desc string) {
	log.Println("os.Args[1]", os.Args[1])
	dingtalk_group_robot.New().SetAccessToken(os.Args[1]).Send(dingtalk_group_robot.NewLink().SetContent(title, desc, "", ""))
}
