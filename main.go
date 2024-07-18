package main

import (
	"flag"
	"owlarticles/conf"

	"owlarticles/router"
)

func main() {
	flag.Parse()
	conf.Init()

	router.Init(conf.Conf)
}
