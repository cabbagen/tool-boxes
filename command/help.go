package command

import "fmt"

type helpItem struct {
  name           string
  description    string
}

var helpItems = []helpItem {
  helpItem { "main version", "get version info" },
  helpItem { "main help", "get help info" },
  helpItem { "main hproxy", "run http proxy" },
}

type help struct {
  name            string
  descriptions    []helpItem
}

func (h help) excute() {
  fmt.Printf("This is tool-boxes params:\n\n")

  for index := 0; index < len(helpItems); index++ {
    fmt.Printf("%s  =>  %s\n", helpItems[index].name, helpItems[index].description)
  }
}

func NewHelp() help {
  return help {
    name: "help",
    descriptions: helpItems,
  }
}
