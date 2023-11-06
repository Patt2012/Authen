package data

import (
	"fmt"
	"log"
	"os"

	"github.com/Patt2012/authen/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Database() (*gorm.DB, error) {

	config, err := configs.InitConfig(".")
	if err != nil {
		log.Fatal("cannot loading config: ", err)
	}
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok", 
			config.DBHost,
			config.DBUser,
			config.DBPassword,
			config.DBName,			
			config.DBPort,			
		)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{	
		Logger: logger.Default.LogMode(logger.Info),	
		DryRun: false,
	})
	if err != nil {
		log.Fatal("Failed to connect database. \n", err)
		os.Exit(2)
	}
	
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	return db, nil
}