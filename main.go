package main

import (
	"github.com/kingcent/supervisor_dingtalk_listener/listener"
)

func main() {
	for {
		listener.Start()
	}
}
