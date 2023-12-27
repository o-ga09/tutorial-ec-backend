package mysql

import (
	"context"
	"log"
	"log/slog"

	"github.com/o-ga09/tutorial-ec-backend/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func New(ctx context.Context) *gorm.DB {
	cfg := config.GetConfig()

	dialector := mysql.Open(cfg.Database_url)
	slog.Info(cfg.Database_url)

	db, err := gorm.Open(dialector, &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: false,
	}}); 

	if err != nil {
		log.Fatal(err)
	}
	db.Logger = db.Logger.LogMode(logger.Silent)
	return db
}