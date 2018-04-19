package plug

import (
	"fmt"
)

var Plugins map[string]Plug

func init() {
	Plugins = make(map[string]Plug)
}

type Plug interface {
	Flag() bool
	Start()
}

func Start() {
	for name, plugin := range Plugins {
		if plugin.Flag() == true {
			go plugin.Start()
			fmt.Printf("Plugin %v starts \n", name)
		} else {
			fmt.Printf("Plugin %v fails \n", name)
		}
	}
}

func Regist(name string, plugin Plug) {
	Plugins[name] = plugin
}
