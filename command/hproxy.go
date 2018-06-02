package command

import "tool-boxes/proxy"

type hproxy struct {
	name           string
	description    string
}

func (hp hproxy) excute() {
  proxy.RunHttpProxy()
}

func NewHproxy() hproxy {
	return hproxy {
		name: "hproxy",
		description: "a http proxy",
	}
} 