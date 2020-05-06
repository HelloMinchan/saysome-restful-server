package controllers

import (
	/* 내부 라이브러리 */
	"database/sql"
	"log"
	"net/http"

	/* 서드파티 라이브러리 */
	"github.com/labstack/echo"

	/* 자체 생성 라이브러리 */
	"github.com/hellominchan/saysome-restful-server/database"
)

// 반환 데이터 json 타입
type userData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// Login 로그인 컨트롤러
func Login(c echo.Context) error {
	// 데이터베이스 연결
	db, err := sql.Open("mysql", database.DB_USER+":"+database.DB_PASSWORD+"@tcp("+database.DB_URL+")/"+database.DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	paramEmail := c.Param("email")
	paramPassword := c.Param("password")

	var email = ""
	var password = ""
	var name = ""

	// 데이터베이스 쿼리 요청
	rows, err := db.Query("SELECT name, email, password FROM member_tb WHERE email = ? and password = ?", paramEmail, paramPassword)
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

	// 데이터베이스 조회 후 값이 들어온 경우
	if email != "" && password != "" && name != "" {
		u := &userData{
			Name:     name,
			Email:    email,
			Password: password,
		}
		return c.JSON(http.StatusOK, u)
	}

	// 데이터베이스 조회 후 아무 것도 없을 경우
	return c.String(http.StatusOK, "Error")
}
