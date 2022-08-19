package rdb

import (
	"fmt"
	"log"

	"entgo.io/ent/examples/fs/ent"
	"github.com/TadayoshiOtsuka/go_test_sample/pkg/config"
	_ "github.com/go-sql-driver/mysql"
)

type SQLHandler struct {
	Client *ent.Client
}

func NewSQLHandler(conf *config.Config) *SQLHandler {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		conf.DB.UserName,
		conf.DB.Password,
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.DBName,
	)

	c, err := ent.Open(conf.DB.Driver, dsn)
	if err != nil {
		log.Panicf("DB Connection Error:%v", err)
	}

	return &SQLHandler{Client: c}
}
