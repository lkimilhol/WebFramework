package config

import (
	"github.com/go-ini/ini"
	"os"
	"path/filepath"
)

type Config struct {
	IP string
	Port string
}

func ReadConfigFile() Config {
	currentDir, err := os.Getwd()
	if err != nil {
		os.Exit(-1)
	}

	configFile, err := ini.Load(filepath.FromSlash(currentDir + "/config/" + "config.ini"))
	if err != nil {
		os.Exit(-2)
	}

	configInfo := Config{}
	section := configFile.Section("Address")
	configInfo.IP = section.Key("IP").String()
	configInfo.Port = section.Key("Port").String()

	if configInfo.IP == "" || configInfo.Port == "" {
		os.Exit(-3)
	}

	return configInfo
}