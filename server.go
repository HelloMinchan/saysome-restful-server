package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_USER     = ""
	DB_PASSWORD = ""
	DB_NAME     = ""
	DB_URL      = ""
)

// db 객체를 echo 라우터 인자로 넘기기위한 타입지정
type (
	dbContext struct {
		echo.Context
		db *sql.DB
	}
)

func main() {
	e := echo.New()

	// 로깅 미들웨어
	e.Use(middleware.Logger())
	// 크래쉬 복구 미들웨어
	e.Use(middleware.Recover())

	// db connect
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASSWORD+"@tcp("+DB_URL+")/"+DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// db 객체를 echo 라우터 인자로 넘기기위한 미들웨어
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &dbContext{c, db}
			return h(cc)
		}
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/users/:email/:password", login)

	// localhost:1323
	e.Logger.Fatal(e.Start(":1323"))
}

func login(c echo.Context) error {
	paramEmail := c.Param("email")
	paramPassword := c.Param("password")

	var email = ""
	var password = ""
	var name = ""

	cd := c.(*dbContext)
	rows, err := cd.db.Query("SELECT name, email, password FROM member_tb WHERE email = ? and password = ?", paramEmail, paramPassword)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&name, &email, &password)
		if err != nil {
			log.Fatal(err)
		}
	}

	if email != "" && password != "" && name != "" {
		return c.String(http.StatusOK, name+" "+email+" "+password)
	}

	return c.String(http.StatusOK, "Error")
}
