package plugin_2

import (
	"fmt"
	"github.com/jeffchanjunwei/plugin-test/plug"
)

func init() {
	p_2 := plugin_2{}
	plug.Regist("plugin_2", p_2)
}

type plugin_2 struct {
}

func (this plugin_2) Flag() bool {
	return true
}

func (this plugin_2) Start() {
	fmt.Println("I am plugin_2")
}
