package database

import (
	"log"

	"github.com/VictorFidelis/learning_gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func ConnectDatabase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Ocorrer uma falha ao conectar ao banco de dados:", err)
	}
	Db.AutoMigrate(&models.Student{})
}
