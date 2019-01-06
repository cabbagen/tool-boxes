package command

import (
	"log"
	"os"
)

type executor interface {
	execute()
}

var commandMap = map[string]executor{
	"version": NewVersion(),
	"help":    NewHelp(),
	"hproxy":  NewHProxy(),
}

func Run() {

	if len(os.Args) != 2 {
		log.Fatalln("please input a command params，look up all command parmas by `main help`")
	}

	name := os.Args[1]

	if commandMap[name] == nil {
		log.Fatalln("input a error command params, look up all command parmas by `main help`")
	}

	commandMap[name].execute()

}
