package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

// Deklarasi Variable Global Untuk Memanggil file Secret Key di Env
var (
	JWTKey = ""
)

type DBConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
	JWTKey string
}

// membuat fungsi global untuk pemanggilan config
func InitConfig() *DBConfig {
	return ReadEnv()
}

func ReadEnv() *DBConfig {
	app := DBConfig{}
	isRead := true

	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.JWTKey = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUser = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPass = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DBPort = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBName = val
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")
		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}
		err = viper.Unmarshal(&app)
		if err != nil {
			log.Println("error parse config : ", err.Error())
			return nil
		}
	}
	JWTKey = app.JWTKey

	return &app
}
