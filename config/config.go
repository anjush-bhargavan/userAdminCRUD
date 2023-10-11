package config

import (
	"os"
	"project_instant/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConfigDB(){
	dsn := os.Getenv("DB_Config")
	var err error

	DB,err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}
	DB.AutoMigrate(&model.Credentials{})

}



func DatabaseMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
			c.Set("db", db)
			c.Next()
	}
}

