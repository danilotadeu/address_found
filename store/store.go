package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/danilotadeu/pismo/store/account"
	"github.com/danilotadeu/pismo/store/transaction"
	_ "github.com/go-sql-driver/mysql"
)

//Container ...
type Container struct {
	Account     account.Store
	Transaction transaction.Store
}

//Register store container
func Register() *Container {
	connectionMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := sql.Open("mysql", connectionMysql)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Println("error db.Ping(): ", err.Error())
		panic(err)
	}

	container := &Container{
		Account:     account.NewStore(db),
		Transaction: transaction.NewStore(db),
	}

	log.Println("Registered -> Store")
	return container
}
