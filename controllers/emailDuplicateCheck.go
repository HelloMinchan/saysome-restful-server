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
type emailDuplicateCheckValue struct {
	CheckValue string `json:"CheckValue"`
}

// EmailDuplicateCheck 이메일 중복체크 컨트롤러
func EmailDuplicateCheck(c echo.Context) error {
	// 데이터베이스 연결
	db, err := sql.Open("mysql", database.DB_USER+":"+database.DB_PASSWORD+"@tcp("+database.DB_URL+")/"+database.DB_NAME)
	// 데이터베이스 서버 에러시 클라이언트에 Database Error 전송
	if err != nil {
		checkValue := &emailDuplicateCheckValue{
			CheckValue: "Database Error",
		}
		return c.JSON(http.StatusOK, checkValue)
	}
	// 데이터베이스 연결 종료 후 닫음
	defer db.Close()

	// 클라이언트에서 넘어온 이메일 데이터
	paramEmail := c.FormValue("email")

	// 만약 이메일 중복 시 회원 이름 임시로 담을 변수
	var nameCol = ""

	/*
		파라미터가 한 개이므로 QueryRow 써도 되지만 그렇게 될 경우 일치하는 데이터가 없을 시
		데이터베이스 내부에서 500에러가 발생하여 404 서버 에러와 혼동되므로 일부러 Query 사용함.
	*/
	// 데이터베이스 회원 정보 쿼리 요청
	rows, err := db.Query("SELECT name_col FROM member_tb WHERE email_col = ?", paramEmail)
	// 데이터베이스 서버 에러시 클라이언트에 Database Error 전송
	if err != nil {
		checkValue := &emailDuplicateCheckValue{
			CheckValue: "Database Error",
		}
		return c.JSON(http.StatusOK, checkValue)
	}
	// 쿼리 종료 후 객체 닫음
	defer rows.Close()

	// 회원 정보 스캔하여 임시 변수에 저장
	for rows.Next() {
		err := rows.Scan(&nameCol)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 해당 이메일이 이미 존재할 경우
	if nameCol != "" {
		checkValue := &emailDuplicateCheckValue{
			CheckValue: "Error",
		}
		return c.JSON(http.StatusOK, checkValue)
	} else {
		// 해당 이메일이 존재하지 않을 경우
		checkValue := &emailDuplicateCheckValue{
			CheckValue: "OK",
		}
		return c.JSON(http.StatusOK, checkValue)
	}
}
