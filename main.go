package main

import (
	_ "SiamLogKit/boot"
	_ "SiamLogKit/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
