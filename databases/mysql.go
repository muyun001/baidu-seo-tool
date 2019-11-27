package databases

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var Db *gorm.DB

func init() {
	var err error

	connArgs := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456",
		"127.0.0.1",
		"3306",
		"rank_task",
	)

	Db, err = gorm.Open("mysql", connArgs)

	if err != nil {
		log.Fatalln(err)
	}

	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetMaxIdleConns(20)
	Db.DB().SetConnMaxLifetime(55 * time.Second)
}
