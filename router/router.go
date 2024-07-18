package router

import "owlarticles/conf"

func Init(c *conf.Config) {
	startHttp(c)
}
