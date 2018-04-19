package plugin_3

import (
	"fmt"
	"github.com/jeffchanjunwei/plugin-test/plug"
)

func init() {
	p_3 := plugin_3{}
	plug.Regist("plugin_3", p_3)
}

type plugin_3 struct {
}

func (this plugin_3) Flag() bool {
	return true
}

func (this plugin_3) Start() {
	fmt.Println("I am plugin_3")
}
