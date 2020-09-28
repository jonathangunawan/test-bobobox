package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jonathangunawan/test-bobobox/entity"
)

type InfraConfigImpl struct {
	cfgFile string
}

func New(cfgFilePath string) InfraConfigImpl {
	return InfraConfigImpl{
		cfgFile: cfgFilePath,
	}
}

func (ic InfraConfigImpl) ReadConfig() (entity.Config, error) {
	var cfg entity.Config

	cfgByte, err := ioutil.ReadFile(ic.cfgFile)
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal([]byte(cfgByte), &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
