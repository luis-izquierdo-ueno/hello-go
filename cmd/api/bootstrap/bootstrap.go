package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"hello-go/internal/creating"
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

	shutdownTimeout = 10 * time.Second
	dbTimeout = 10 * time.Second
)


func Run() error {
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPassword, dbHost, dbPort, dbName )
	db, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db, dbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository)
	ctx, srv := server.New(context.Background(), host,port, shutdownTimeout, *creatingCourseService)
	
	return srv.Run(ctx)
}