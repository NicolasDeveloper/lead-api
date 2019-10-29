package env

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	Parameters *Configuration
)

type Configuration struct {
	Database Database
}

type Database struct {
	Ip       string
	Port     string
	AppPort  string
	Username string
	Password string
}

func InitConfiguration() {
	switch os.Getenv("ENVIRONMENT") {
	default:
		viper.SetConfigName("prod.env")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./env")
		viper.ReadInConfig()
	}

	err := viper.Unmarshal(&Parameters)

	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
}
	err := viper.Unmarshal(&Parameters)

	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
}

	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}
}
