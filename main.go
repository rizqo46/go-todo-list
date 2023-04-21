package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/bytedance/sonic"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/rizqo46/go-todo-list/db"
	"github.com/rizqo46/go-todo-list/repo"
	"github.com/rizqo46/go-todo-list/route"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(repo.Todo), new(repo.Activity))
	orm.RegisterDataBase("default", "mysql", getDbUrl())
	orm.SetDataBaseTZ("default", time.UTC)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	o := orm.NewOrm()
	db.Migrate(o)

	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	route.InitRoute(app, o)

	app.Listen(":3030")
}

func getDbUrl() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DBNAME"),
	)
}
