package internal

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	Addr       string `json:"addr"`
	RoboMaster struct {
		Address string `json:"address"`
	} `json:"rm"`
}

var (
	Config *AppConfig
	DEBUG  = os.Getenv("DEBUG") == "true"
)

func bootConfig() {
	f, err := os.Open("config.json")
	if err != nil {
		f, err = os.OpenFile("config.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to open config file")
		}

		defaultConfig, er := sonic.ConfigFastest.MarshalIndent(&AppConfig{}, "", "  ")
		if er == nil {
			f.Write(defaultConfig)
		}

		log.Fatal().Err(err).Msg("Failed to open config file")
	}

	defer f.Close()

	Config = new(AppConfig)

	if err = sonic.ConfigFastest.NewDecoder(f).Decode(Config); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config file")
	}

	log.Info().Interface("config", Config).Msg("AppConfig loaded")

	if DEBUG {
		gin.SetMode(gin.DebugMode)
		log.Info().Msg("Debug mode enabled")
	} else {
		gin.SetMode(gin.ReleaseMode)
		log.Info().Msg("Release mode enabled")
	}
}
