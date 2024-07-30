package internal

import (
	"os"

	"github.com/bytedance/sonic"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Port       string `json:"port"`
	RoboMaster struct {
		Address string
	} `json:"rm"`
}

var (
	config *Config
	DEBUG  = os.Getenv("DEBUG") == "true"
)

func bootConfig() {
	f, err := os.Open("config.json")
	if err != nil {
		f, err = os.OpenFile("config.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to open config file")
		}

		defaultConfig, er := sonic.ConfigFastest.MarshalIndent(&Config{}, "", "  ")
		if er == nil {
			f.Write(defaultConfig)
		}

		log.Fatal().Err(err).Msg("Failed to open config file")
	}

	defer f.Close()

	if err = sonic.ConfigFastest.NewDecoder(f).Decode(config); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config file")
	}

	log.Info().Interface("config", config).Msg("Config loaded")
}
