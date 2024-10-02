package app

import (
	"github.com/nil-go/konf"
	"github.com/nil-go/konf/provider/file"
)

type Config struct {
	Logging struct {
		Level int
	}
	Web WebConfig
	DB  struct {
		DriverName       string
		ConnectionString string
	}
}

func ReadLocalConfig(configPath string) (Config, error) {
	var config konf.Config

	err := config.Load(file.New(configPath))
	if err != nil {
		return Config{}, err
	}

	var res Config

	err = config.Unmarshal("", &res)
	if err != nil {
		return Config{}, err
	}

	return res, nil
}
