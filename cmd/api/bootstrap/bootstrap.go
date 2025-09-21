package bootstrap

import (
	"database/sql"
	"fmt"

	"hello-go/internal/platform/server"
	"hello-go/internal/platform/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "root"
	dbPassword = "root"
	dbHost = "localhost"
	dbPort = 3306
	dbName = "hello-go"
)


func Run() error {
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName )
	db, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)
	srv := server.New(host,port, courseRepository)
	
	return srv.Run()
}