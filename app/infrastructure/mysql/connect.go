package mysql

import (
	"context"
	"log"

	"github.com/o-ga09/tutorial-go-fr/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func New(ctx context.Context) *gorm.DB {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	dialector := mysql.Open(cfg.Database_url)

	db, err := gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: false,
	}}); 

	if err != nil {
		log.Fatal(err)
	}
	db.Logger = db.Logger.LogMode(logger.Silent)
	return db
}