package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shahid4183/Golang/REST/models"
	"github.com/Shahid4183/Golang/REST/server"
	"github.com/Shahid4183/Golang/REST/utility"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	config "github.com/spf13/viper"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	log.Println("GET CONFIG")

	config.AddConfigPath(".")
	config.SetConfigName("config")

	if err := config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error getting config from file: %s", err))
	}
	logger := lumberjack.Logger{
		Filename:   config.GetString("logging.directory") + config.GetString("logging.filename"),
		MaxSize:    config.GetInt("logging.size"), // megabytes
		MaxBackups: 1,
		MaxAge:     1,    //days
		Compress:   true, // disabled by default
	}
	log.SetOutput(&logger)

	log.Println("CONNECT TO DATABASE")
	database, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s sslrootcert=%s",
			config.GetString("db.host"),
			config.GetString("db.port"),
			config.GetString("db.user"),
			config.GetString("db.name"),
			config.GetString("db.password"),
			utility.IfThenElse(config.GetString("env") == "local", "disable", config.GetString("db.sslmode")),
			config.GetString("db.sslrootcert"),
		),
	)

	if err != nil {
		panic("COULDN'T CONNECT TO DATABASE " + err.Error())
	}
	database.SetLogger(log.New(&logger, "\n\n", 0))
	if len(os.Args) > 1 && (os.Args[1] == "-automigrate" || os.Args[1] == "-a") {
		database.AutoMigrate(
			&models.User{},
		)
	}
	log.Println("START UP SERVER")
	apiServer, err := server.New(database, &logger)
	if err != nil {
		panic(fmt.Errorf("fatal error setting up server: %s", err))
	}
	port := ":" + config.GetString("server.port")
	if config.GetString("env") != "local" {
		apiServer.Logger.Fatal(apiServer.StartTLS(port, config.GetString("server.sslCert"), config.GetString("server.sslKey")))
	} else {
		apiServer.Logger.Fatal(apiServer.Start(port))
	}
}
