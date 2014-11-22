package models

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Table interface {}

type TableService interface {
    Create() error
    Retrieve(id int) (interface{}, error)
    Update() error
    Delete(id int) error
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// Initialize gorp for struct mapping
func InitDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "/tmp/gorptest.sqlite")
	checkErr(err, "DB INIT")

	err = db.Ping()
	checkErr(err, "DB PING")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.TraceOn("[gorp]", log.New(os.Stdout, "myapp:", log.Lmicroseconds))

	dbmap.AddTableWithName(User{}, "users").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

    return dbmap
}

// Initalize Services
func InitTableServices(dbmap *gorp.DbMap) map[string]TableService {
    services := make(map[string]TableService)
    _ = services
    services["user"] = NewUserService(dbmap)

	return services
}


