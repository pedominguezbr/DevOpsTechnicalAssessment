package config

import "github.com/jinzhu/configor"

type Config struct {
	AppName string
	Port    string
	DB      struct {
		// the follows databases were injected for Hygen
		Oracle []struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"localhost"`
			Port     string `default:"5532"`
			Username string `default:"freyja"`
			Password string `default:"Spsa2700"`
			Database string `default:"pmm"`
			Nameconn string `default:"pmm"`
		}
	}
	Email struct {
		Id  string
		Url string
	}
	Contacts struct {
		Name  string
		Email string
	}
}

func NewConfig() (*Config, error) {
	c := &Config{}
	err := configor.Load(c, "config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
