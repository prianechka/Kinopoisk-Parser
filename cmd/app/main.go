package main

import (
	"Kinopoisk-Parser/config"
	"Kinopoisk-Parser/internal/server"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"os"
)

var configPath = os.Getenv("CONFIG_FILE")

func main() {
	appConfig := config.CreateConfig()
	_, err := toml.DecodeFile(configPath, &appConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetReportCaller(false)
	logrus.SetFormatter(&logrus.TextFormatter{PadLevelText: false, DisableLevelTruncation: false})

	appServer := server.CreateServer(appConfig)

	err = appServer.Start()
	if err != nil {
		panic(err)
	}
}
