package command

import "os"
import "log"

type executor interface {
	excute()
}

var commandMap = map[string]executor{
	"version": NewVersion(),
	"help": NewHelp(),
	"hproxy": NewHproxy(),
}

func Run() {
	
	if len(os.Args) != 2 {
		log.Fatalln("please input a command params，look up all command parmas by `main help`")
	}

	name := os.Args[1]

	if commandMap[name] == nil {
		log.Fatalln("input a error command params, look up all command parmas by `main help`")
	}

	commandMap[name].excute()

}