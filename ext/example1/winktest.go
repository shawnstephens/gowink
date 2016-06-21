package main

import (
	"os"

	"github.com/dan9186/gowink"
	log "github.com/gomicro/ledger"
	"github.com/kelseyhightower/envconfig"
)

const (
	app = "WINK"
)

type Configuration struct {
	ClientId     string `required:"true"`
	ClientSecret string `required:"true"`
	User         string `required:"true"`
	Password     string `required:"true"`

	WinkAPI string `default:"https://api.wink.com/"`
}

var (
	config Configuration
)

func main() {
	err := envconfig.Process(app, &config)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	wink := gowink.New(config.WinkAPI)

	err = wink.SignIn(config.ClientId, config.ClientSecret, config.User, config.Password)
	if err != nil {
		log.Error(err.Error())
		return
	}

	b, err := wink.GetDevices()
	if err != nil {
		log.Error(err.Error())
		return
	}

	log.Debug(string(b))
}
