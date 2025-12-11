package configPkg

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	Env  *Config
	once sync.Once
)

func Init() {

	once.Do(func() {

		viper.SetConfigName("config")
		viper.SetConfigType("yml")
		viper.AddConfigPath("../../")

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&Env); err != nil {
			panic(err)
		}

		fmt.Println(Env)
	})

}
