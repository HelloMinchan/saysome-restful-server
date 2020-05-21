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
type LogincheckValue struct {
	CheckValue string `json:"CheckValue"`
}

// 반환 데이터 json 타입
type userData struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	Name                 string `json:"name"`
	Food                 string `json:"food"`
	ProvisionAccept      string `json:"provisionAccept"`
	EmailReceptionAccept string `json:"emailReceptionAccept"`
}

// Login 로그인 컨트롤러
func Login(c echo.Context) error {
	// 데이터베이스 연결
	db, err := sql.Open("mysql", database.DB_USER+":"+database.DB_PASSWORD+"@tcp("+database.DB_URL+")/"+database.DB_NAME)
	// 데이터베이스 서버 에러시 클라이언트에 Database Error 전송
	if err != nil {
		checkValue := &LogincheckValue{
			CheckValue: "Database Error",
		}
		return c.JSON(http.StatusOK, checkValue)
	}
	// 데이터베이스 연결 종료 후 닫음
	defer db.Close()

	// 클라이언트에서 넘어온 이메일 데이터
	paramEmail := c.FormValue("email")
	// 클라이언트에서 넘어온 비밀번호 데이터
	paramPassword := c.FormValue("password")

	// 회원정보 임시로 담을 변수
	var emailCol = ""
	var passwordCOl = ""
	var nameCol = ""
	var foodCol = ""
	var provisionAcceptCol = ""
	var emailReceptionAcceptCol = ""

	// 데이터베이스 회원 정보 쿼리 요청
	rows, err := db.Query("SELECT email_col, password_col, name_col, food_col, provisionAccept_col, emailReceptionAccept_col FROM member_tb WHERE email_col = ? and password_col = ?", paramEmail, paramPassword)
	// 데이터베이스 서버 에러시 클라이언트에 Database Error 전송
	if err != nil {
		checkValue := &LogincheckValue{
			CheckValue: "Database Error",
		}
		return c.JSON(http.StatusOK, checkValue)
	}
	// 쿼리 종료 후 객체 닫음
	defer rows.Close()

	// 회원 정보 다중 컬럼 스캔하여 임시 변수에 저장
	for rows.Next() {
		err := rows.Scan(&emailCol, &passwordCOl, &nameCol, &foodCol, &provisionAcceptCol, &emailReceptionAcceptCol)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 데이터베이스 조회 후 값이 들어온 경우
	if emailCol != "" && passwordCOl != "" && nameCol != "" && foodCol != "" {
		u := &userData{
			Email:                emailCol,
			Password:             passwordCOl,
			Name:                 nameCol,
			Food:                 foodCol,
			ProvisionAccept:      provisionAcceptCol,
			EmailReceptionAccept: emailReceptionAcceptCol,
		}
		return c.JSON(http.StatusOK, u)
	}

	// 데이터베이스 조회 후 아무 것도 없을 경우
	return c.String(http.StatusOK, "Error")
}
