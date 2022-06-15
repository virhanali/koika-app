package database

import (
	"github.com/sirupsen/logrus"
	"github.com/virhanali/koika-app/config/pkg"
	"github.com/virhanali/koika-app/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(pkg.Godotenv("PG_URL")), &gorm.Config{})

	if err != nil {
		defer logrus.Info("connection database failed")
		logrus.Fatal(err)
		return nil
	}
	logrus.Info("database connection success")

	err = db.AutoMigrate(&models.ProductModel{}, &models.OutletModel{}, &models.SupplierModel{}, &models.TransactionModel{}, &models.MerchantModel{}, &models.CustomerModel{}, &models.RoleModel{}, &models.UserModel{})
	if err != nil {
		defer logrus.Info("Database migration failed")
		logrus.Fatal(err)
		return nil
	}

	return db
}
