package mysql

import (
	"context"
	"log/slog"
	"time"

	"github.com/o-ga09/tutorial-ec-backend/app/config"
	model "github.com/o-ga09/tutorial-ec-backend/app/infrastructure/mysql/schema"
	"github.com/o-ga09/tutorial-ec-backend/app/server/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func New(ctx context.Context) *gorm.DB {
	cfg := config.GetConfig()

	dialector := mysql.Open(cfg.Database_url)

	connect(dialector,100)
	slog.Log(context.Background(), middleware.SeverityInfo, "db connect")

	if cfg.Env == "dev" {
		db.AutoMigrate(model.User{},model.Product{},model.Owner{},model.OrderProduct{},model.Order{})
	}
	db.Logger = db.Logger.LogMode(logger.Silent)
	return db
}

func connect(dialector gorm.Dialector, count uint) {
	var err error
	if db, err = gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // AutoMigrateで外部キー制約にしない
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,  // テーブル名を単数系にしない
	}});  err != nil {
		if count > 1 {
			time.Sleep(time.Second * 10)
			count--
			slog.Log(context.Background(), middleware.SeverityInfo, "db connection retry")
			connect(dialector, count)
			return
		}
		slog.Log(context.Background(), middleware.SeverityInfo, "db connection retry count 100")
		panic(err.Error())
	}
}