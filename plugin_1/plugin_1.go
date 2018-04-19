package plugin_1

import (
	"fmt"
	"github.com/jeffchanjunwei/plugin-test/plug"
)

func init() {
	p_1 := plugin_1{}
	plug.Regist("plugin_1", p_1)
}

type plugin_1 struct {
}

func (this plugin_1) Flag() bool {
	return true
}

func (this plugin_1) Start() {
	fmt.Println("I am plugin_1")
}
