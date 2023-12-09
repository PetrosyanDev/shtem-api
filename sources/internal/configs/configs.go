// Erik Petrosyan ©
package configs

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Configs struct {
	TLS    bool
	Global struct {
		CDNBaseURL   string `json:"cdnBaseUrl"`
		PublicAPIKey string `json:"publicApiKey"`
	} `json:"global"`
	API struct {
		APIServiceKey string `json:"apiServiceKey"`
		HeaderKey     string `json:"headerKey"`
		Addr          string `json:"address"`
	} `json:"api"`
	PostgresDB struct {
		Address string `json:"dbAddress"`
		Port    string `json:"dbPort"`
		DB      string `json:"db"`
		User    string `json:"dbUser"`
		Pass    string `json:"dbPass"`
		// TLS     bool   `json:"tls"`
	} `json:"postgres"`
}

func NewConfigs(args []string) (*Configs, error) {
	cfgPath := ""
	cfg := new(Configs)
	for i, arg := range args {
		if arg == "--tls" {
			cfg.TLS = true
		} else if arg == "--cfg" {
			log.Println("--cfg", args[i+1])
			if len(args) > i {
				cfgPath = args[i+1]
			}
		}
	}
	if cfgPath == "" {
		return nil, errors.New("config file does not specified")
	}
	if data, err := os.ReadFile(cfgPath); err == nil {
		if err = json.Unmarshal(data, cfg); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}
	return cfg, nil
}
