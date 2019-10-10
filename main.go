package main

import (
	"github.com/irebit/supervisor_dingtalk_listener/listener"
)

func main() {
	for {
		listener.Start()
	}
}
