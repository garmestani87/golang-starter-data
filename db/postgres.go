package db

import (
	"fmt"
	"golang-starter-data/config"
	"golang-starter-data/domain/entity"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbc *gorm.DB

func InitPostgres(cfg *config.Config) {
	conn := "postgresql://" + cfg.Data.User + ":" + cfg.Data.Password + "@" + cfg.Data.Host + "/" + cfg.Data.DbName + "?sslmode=" + cfg.Data.SslMode

	fmt.Printf("connection string is : %s", conn)

	var err error
	dbc, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDb, err := dbc.DB()
	if err != nil {
		panic(err)
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(cfg.Postgres.ConnMaxLifetime) * time.Minute)

	err = dbc.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
}

func GetDatabase() *gorm.DB {
	return dbc
}

func CloseDatabase() {
	conn, _ := dbc.DB()
	err := conn.Close()
	if err != nil {
		panic(err)
	}
}
