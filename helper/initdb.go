package helper

import (
	"fmt"

	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func GetDbini() (config.Configer, error) {
	dbconf, err := config.NewConfig("ini", "config/db.ini")
	if err != nil {
		fmt.Println("ini failed,", err)
		return nil, err
	}
	return dbconf, nil
}

type DbSetting struct {
	Localhost string
	Port      int
	User      string
	Pwd       string
	Table     string
}

func init() {
	dbini, _ := GetDbini()
	localhost := dbini.String("db_read::DBHOST")
	port, _ := dbini.Int("db_read::DBPORT")
	dbuser := dbini.String("db_read::DBUSER")
	dbpwd := dbini.String("db_read::DBPASSWORD")
	dbtable := dbini.String("db_read::DBTABLE")
	sets := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbuser, dbpwd, localhost, port, dbtable)
	fmt.Println("sss ", sets)
	database, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbuser, dbpwd, localhost, port, dbtable))

	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	fmt.Println("ddd")
}
