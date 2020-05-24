# saysome-restful-server
SaySome RESTful API Server.
- - -
* 흐름도
<img src="https://user-images.githubusercontent.com/52199223/82762655-348beb00-9e3d-11ea-8694-f72eacdfffb2.PNG"><img>   
- - -
* 빌드 버전 기록   
  * Go : 1.14.3
- - -
* 의존 프레임워크 및 라이브러리
  * Echo : go get -u github.com/labstack/echo/...
  * go-sql-driver : go get -u github.com/go-sql-driver/mysql
- - -
* clone 후 해야 할 일
  * /database/config.go 초기화 (파일명 및 변수명에서 fake_ 지울 것)
- - -
* RESFful API
  * /login : 로그인 API
    + 코드 위치 : /controllers/login.go
    + HTTP request methods : POST
    + Input : email && password
    + Output : userData || "Database Error" || "Error"
  * /emailDuplicatecheck : 이메일 중복체크 API
    + 코드 위치 : /controllers/emailDuplicateCheck.go
    + HTTP request methods : POST
    + Input : email
    + Output : "OK" || "Database Error" || "Error"
  * /signupapply : 회원가입 신청 API
    + 코드 위치 : /controllers/signupApply.go
    + HTTP request methods : POST
    + Input : email && password && name && food && provisionAccept && emailReceptionAccept
    + Output : "OK" || "Database Error" || "Database Error2"
