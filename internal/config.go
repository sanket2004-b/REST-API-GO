package config

import (
	"flag"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address"`
}
type Config struct {
	Env         string `yaml:"env" env:"ENV" required:"true" envDefault:"production"`
	StoragePath string `yaml:"storage_path"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad() *config {
	var configPath string
	if configPath == "" {
		flags := flag.String("config", "", "path to config file")
		flags.Parse()
		configPath=*flags
		if cofigPath==""{
			log.Fatal("Config path is not set")
		}
	}
	var cfg Config
	err:=cleanenv.ReadConfig(configPath,&cfg){
		if err != nil{
			log.Fatalf("can not read config file : %s",err.Error())
		}

		return &cfg
	}
}