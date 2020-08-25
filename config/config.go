package config

import (
	"gopkg.in/ini.v1"
)

type Config struct {
	AsfLogPath string
}

func GetConfig(filename string) Config {
	cfg, err := ini.Load(filename)
	if err != nil {
		panic(err)
	}
	section, err := cfg.GetSection("")
	if err != nil {
		panic(err)
	}
	logPath, err := section.GetKey("asf-log")
	if err != nil {
		panic(err)
	}

	return Config{AsfLogPath: logPath.String()}
}
