package main

import (
	/* 내부 라이브러리 */
	"net/http"

	/* 서드파티 라이브러리 */
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"

	/* 자체 생성 라이브러리 */
	"github.com/hellominchan/saysome-restful-server/controllers"
)

func main() {
	e := echo.New()

	// 로깅 미들웨어
	e.Use(middleware.Logger())
	// 크래쉬 복구 미들웨어
	e.Use(middleware.Recover())

	// CORS 미들웨어
	e.Use(middleware.CORS())

	// 메인 API
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	// 로그인 API
	e.GET("/login/:email/:password", controllers.Login)

	// 서버 생성 Port 1323
	e.Logger.Fatal(e.Start(":1323"))
}
