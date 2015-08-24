package gibson

import (
	"fmt"
	"gibson/utils"
	"github.com/spf13/viper"
	"reflect"
)

// TODO: add application config
var (
	Configs map[string]string = map[string]string{
		"routes": "conf/",
	}
	config *ConfigStruct
)

type ConfigStruct struct {
	Routes *viper.Viper
}

func GetConfigInstance() *ConfigStruct {
	if config == nil {
		config = newConfig()
	}
	return config
}

func newConfig() *ConfigStruct {
	config := &ConfigStruct{}
	e := reflect.ValueOf(config).Elem()
	for kind, path := range Configs {
		v := reflect.ValueOf(createConfigObject(kind, path))
		e.FieldByName(utils.Capitalize(kind)).Set(v)
	}

	return config
}

func createConfigObject(kind, path string) *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(kind)
	v.AddConfigPath(path)

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return v
}
