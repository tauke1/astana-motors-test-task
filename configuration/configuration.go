package configuration

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type configuration struct {
	RedisAddress       string
	RedisPassword      string
	RedisDatabase      int
	JwtSecret          string
	JwtIssuer          string
	JwtExpirationHours int
	DbHost             string
	DbPassword         string
	DbUser             string
	DbName             string
	ServerPort         int
}

var C configuration

func LoadConfig(path string) {
	if path == "" {
		panic("path argument must not empty")
	}

	Config := &C
	viper.SetConfigType("yml")
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
