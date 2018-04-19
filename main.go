package main

import (
	"github.com/jeffchanjunwei/plugin-test/plug"
	_ "github.com/jeffchanjunwei/plugin-test/plugin_1"
	_ "github.com/jeffchanjunwei/plugin-test/plugin_2"
	_ "github.com/jeffchanjunwei/plugin-test/plugin_3"
	"time"
)

func main() {
	plug.Start()
	time.Sleep(5 * time.Second)
}
