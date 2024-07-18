package conf

import (
	"flag"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/quan-xie/tuba/database/mongo"
)

var (
	Conf     = &Config{}
	confPath string
)

type Config struct {
	Env   string
	Mongo *mongo.Config
}

func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

// Init config from config file.
func Init() {
	// TODO 从配置中心获取配置文件
	deployEnv := os.Getenv("DEPLOY_ENV")

	if deployEnv == "" {
		deployEnv = "local"
	}
	switch deployEnv {
	case "local":
		confPath = "./conf/local.toml"
	case "develop":
		confPath = "./conf/dev.toml"
	case "qa":
		confPath = "./conf/qa.toml"
	case "test":
		confPath = "./conf/test.toml"
	case "staging":
		confPath = "./conf/staging.toml"
	case "prod":
		confPath = "./conf/prod.toml"
	}
	if _, err := toml.DecodeFile(confPath, &Conf); err != nil {
		panic(err)
	}
	Conf.Env = deployEnv
}
