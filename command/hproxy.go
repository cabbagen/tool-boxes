package command

import "tool-boxes/proxy"

type hproxy struct {
	name        string
	description string
}

func (hp hproxy) execute() {
	proxy.RunHttpProxy()
}

func NewHProxy() hproxy {
	return hproxy{
		name:        "hproxy",
		description: "a http proxy",
	}
}
