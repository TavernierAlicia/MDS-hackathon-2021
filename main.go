package main

import (
	"go.uber.org/zap"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// init zap
var (
	log *zap.Logger
)

func main() {

	log, _ = zap.NewProduction()
	defer log.Sync()

	Router()
}

func RunDb() (db *sqlx.DB, err error) {

	//// IMPORT CONFIG ////
	viper.SetConfigName("conf")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Error("Unable to load config file")
	}

	//// DEFINE CONFIG VARIABLES FROM JSON FILE ////
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.pass")
	dbname := viper.GetString("database.dbname")

	//// DB CONNECTION ////
	pathSQL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, pass, host, port, dbname)
	db, err = sqlx.Connect("mysql", pathSQL)

	if err != nil {
		log.Error("unmarshal error")
	}
	return db, err
}