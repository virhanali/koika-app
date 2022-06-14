package database

import (
	"github.com/sirupsen/logrus"
	"github.com/virhanali/koika-app/config/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open(pkg.Godotenv("MYSQL_URL")), &gorm.Config{})

	if err != nil {
		defer logrus.Info("connection database failed")
		logrus.Fatal(err)
		return nil
	}
	logrus.Info("database connection success")
	return db
}
