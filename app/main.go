package main

import (
	"CleanTodo/todo/delivery"
	"CleanTodo/todo/repository/cache"
	"CleanTodo/todo/repository/mysql"
	"CleanTodo/todo/usecase"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)

const (
	//for mysql
	UserName     string = "root"
	Password     string = "chungyoMarc"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "todo"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
	//for redis
	Cachemethod   string = "tcp"
	CachemAddress string = "127.0.0.1:6379"
)

func main() {
	//redis
	cacheConn, err := redis.Dial(Cachemethod, CachemAddress)

	//mysql
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", UserName, Password, Addr, Port, Database)
	dbConn, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	g := gin.Default()

	todoRepo := mysql.NewMysqlTodoRepository(dbConn)
	todoCache := cache.NewRedisTodoCache(cacheConn)
	todoUsecase := usecase.NewTodoUsecase(todoRepo, todoCache)

	delivery.NewTodoRouter(g, todoUsecase)
}
