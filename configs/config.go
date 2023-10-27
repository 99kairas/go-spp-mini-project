package configs

import (
	"fmt"
	"go-spp/models"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc

	config := Config{
		DB_Username: os.Getenv("DB_Username"),
		DB_Password: os.Getenv("DB_Password"),
		DB_Port:     os.Getenv("DB_Port"),
		DB_Host:     os.Getenv("DB_Host"),
		DB_Name:     os.Getenv("DB_Name"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	InitMigrate()
	Seeders()

	return DB
}

func InitMigrate() {
	err := DB.AutoMigrate(&models.Admin{}, &models.Student{}, &models.Grade{}, &models.SPP{}, &models.Payment{})
	if err != nil {
		panic("failed to migrate database")
	}

}

func Seeders() {
	grade := []models.Grade{
		{
			ID:    uuid.New(),
			Level: "First Grade",
		},
		{
			ID:    uuid.New(),
			Level: "Second Grade",
		},
		{
			ID:    uuid.New(),
			Level: "Third Grade",
		},
		{
			ID:    uuid.New(),
			Level: "Fourth Grade",
		},
		{
			ID:    uuid.New(),
			Level: "Fifth Grade",
		},
		{
			ID:    uuid.New(),
			Level: "Sixth Grade",
		},
	}

	for _, v := range grade {
		var exist models.Grade

		err := DB.Where("level = ?", v.Level).First(&exist).Error

		if err != nil {
			DB.Create(&v)
		}
	}
}
