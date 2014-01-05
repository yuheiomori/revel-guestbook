package controllers

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	r "github.com/robfig/revel"

	"net/url"
	"os"
	"revel-guestbook/app/models"
)

var (
	Dbm *gorp.DbMap
)

// データソース文字列を変換
func convert_datasource(ds string) (result string) {
	url, _ := url.Parse(ds)
	result = fmt.Sprintf("%s@tcp(%s:3306)%s", url.User.String(), url.Host, url.Path)
	return
}

func Init() {
	var datasource string
	// for heroku with cleardb
	if os.Getenv("CLEARDB_DATABASE_URL") != "" {
		datasource = convert_datasource(os.Getenv("CLEARDB_DATABASE_URL"))
	} else {
		datasource = "user:pass@/database_name?charset=utf8"
	}
	db, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}
	Dbm = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	Dbm.AddTableWithName(models.Greeting{}, "greetings").SetKeys(true, "Id")

	Dbm.TraceOn("[gorp]", r.INFO)
	err = Dbm.CreateTablesIfNotExists()
	if err != nil {
		panic(err)
	}

}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
