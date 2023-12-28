package mysql

import (
	"context"

	"github.com/o-ga09/tutorial-ec-backend/app/config"
	model "github.com/o-ga09/tutorial-ec-backend/app/infrastructure/mysql/schema"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func New(ctx context.Context) *gorm.DB {
	cfg := config.GetConfig()

	dialector := mysql.Open(cfg.Database_url)

	db, err := gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // AutoMigrateで外部キー制約にしない
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,  // テーブル名を単数系にしない
	}}); 

	if err != nil {
		//panic(err)
	}

	if cfg.Env == "dev" {
		db.AutoMigrate(model.User{},model.Product{},model.Owner{},model.OrderProduct{},model.Order{})
	}
	db.Logger = db.Logger.LogMode(logger.Silent)
	return db
}