package main

import "supervisor_dingtalk_listener/listener"

func main() {
	for {
		listener.Start()
	}
}
