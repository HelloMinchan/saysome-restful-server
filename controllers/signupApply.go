package controllers

import (
	/* 내부 라이브러리 */
	"database/sql"
	"net/http"
	"strconv"

	/* 서드파티 라이브러리 */
	"github.com/labstack/echo"

	/* 자체 생성 라이브러리 */
	"github.com/hellominchan/saysome-restful-server/database"
)

// 반환 데이터 json 타입
type signupApplyValue struct {
	CheckValue string `json:"CheckValue"`
}

// SignupApply 회원가입 신청 컨트롤러
func SignupApply(c echo.Context) error {
	// 데이터베이스 연결
	db, err := sql.Open("mysql", database.DB_USER+":"+database.DB_PASSWORD+"@tcp("+database.DB_URL+")/"+database.DB_NAME)
	// 데이터베이스 서버 에러시 클라이언트에 Database Error 전송
	if err != nil {
		checkValue := &signupApplyValue{
			CheckValue: "Database Error",
		}
		return c.JSON(http.StatusOK, checkValue)
	}
	// 데이터베이스 연결 종료 후 닫음
	defer db.Close()

	// 클라이언트에서 넘어온 이메일 데이터
	paramEmail := c.FormValue("email")
	paramPassword := c.FormValue("password")
	paramName := c.FormValue("name")
	paramFood := c.FormValue("food")
	paramProvisionAccept := c.FormValue("provisionAccept")
	paramEmailReceptionAccept := c.FormValue("emailReceptionAccept")

	// 이용약관 동의 string으로 넘어온 값 boolean 타입으로 다시 형변환
	convertedParamProvisionAccept, err := strconv.ParseBool(paramProvisionAccept)
	// 이메일 수신 동의 string으로 넘어온 값 boolean 타입으로 다시 형변환
	convertedparamEmailReceptionAccept, err := strconv.ParseBool(paramEmailReceptionAccept)

	// 데이터베이스 회원 정보 기록 쿼리 요청
	result, err := db.Exec("INSERT INTO member_tb VALUES (?, ?, ?, ? ,? ,?)", paramEmail, paramPassword, paramName, paramFood, convertedParamProvisionAccept, convertedparamEmailReceptionAccept)
	// 데이터베이스 서버 에러시 클라이언트에 Database Error 전송
	if err != nil {
		checkValue := &signupApplyValue{
			CheckValue: "Database Error",
		}
		return c.JSON(http.StatusOK, checkValue)
	}

	// 쿼리 결과 반환
	n, err := result.RowsAffected()

	if n == 1 {
		// 1줄의 Row가 기록 된 경우(기록 성공)
		checkValue := &signupApplyValue{
			CheckValue: "OK",
		}
		return c.JSON(http.StatusOK, checkValue)
	} else {
		// 그렇지 않은 경우(기록 실패)
		checkValue := &signupApplyValue{
			CheckValue: "Database Error2",
		}
		return c.JSON(http.StatusOK, checkValue)
	}
}
