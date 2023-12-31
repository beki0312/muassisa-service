package db

import (
	"fmt"
	"go.uber.org/fx"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"muassisa-service/internal/config"
	"muassisa-service/internal/pkg/logger"
)

var NewPostgres = fx.Provide(newPostgres)

type dependencies struct {
	fx.In
	Config config.IConfig
	Logger logger.ILogger
}

type postgres struct {
	Postgres *gorm.DB
}

type IPostgres interface {
	GetPostgresConnection() *gorm.DB
}

func newPostgres(dp dependencies) IPostgres {

	host := dp.Config.GetString("api.postgres.host")
	port := dp.Config.GetString("api.postgres.port")
	user := dp.Config.GetString("api.postgres.user")
	dbname := dp.Config.GetString("api.postgres.dbname")
	password := dp.Config.GetString("api.postgres.password")

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dushanbe",
		host, user, password, dbname, port)
	conn, err := gorm.Open(postgresDriver.Open(connString))

	if err != nil {
		dp.Logger.Error("%s", "GetPostgresConnection -> Open error: ", err.Error())
		return nil
	}

	log.Println("Postgres connection success: ", host)
	dp.Logger.Info("%s", "Postgres connection success: ", host)
	return &postgres{Postgres: conn}
}

func (p *postgres) GetPostgresConnection() *gorm.DB {
	return p.Postgres
}
