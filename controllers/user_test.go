package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Fishmansky/noteflow/inits"
	"github.com/Fishmansky/noteflow/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginSuite struct {
	suite.Suite
	DB *gorm.DB
}

func TestLoginSuite(t *testing.T) {
	suite.Run(t, &LoginSuite{})
}

func (s *LoginSuite) SetupSuite() {
	if err := godotenv.Load("../.env"); err != nil {
		s.FailNow("Error loading .env file")
	}
	inits.ConnectToDB()
	inits.SyncDB()
	inits.ConnecRedis()
	s.DB = inits.DB
	bytes, err := bcrypt.GenerateFromPassword([]byte("Test1234"), 14)
	if err != nil {
		s.FailNow("Cloudn't hash password for test user")
	}
	s.DB.Create(&models.User{Email: "existing@user.pl", Password: string(bytes)})
	fmt.Println("Test user created!")
}

func (s *LoginSuite) TearDownSuite() {
	s.DB.Where("email LIKE ?", "existing@user.pl").Delete(&models.User{})
	fmt.Println("Test user deleted!")
}

func GetTestGinContext(w *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	return ctx
}

func MockJsonPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

type loginTestData struct {
	Email    string
	Password string
}

var loginTestDataTable = []loginTestData{
	{"existing@user.pl", "Test1234"},
	{"nonexisting@user.pl", "Test1234"},
	{"", ""},
}

func (s *LoginSuite) TestLoginExistingUser() {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonPost(ctx, loginTestDataTable[0])
	Login(ctx)
	s.EqualValues(http.StatusOK, w.Code)
}

func (s *LoginSuite) TestLoginNonexistingUser() {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonPost(ctx, loginTestDataTable[1])
	Login(ctx)
	s.EqualValues(http.StatusBadRequest, w.Code)
}

func (s *LoginSuite) TestLoginEmptyCredentials() {
	w := httptest.NewRecorder()
	ctx := GetTestGinContext(w)
	MockJsonPost(ctx, loginTestDataTable[2])
	Login(ctx)
	s.EqualValues(http.StatusBadRequest, w.Code)
}

// func (lt *LoginSuite) TestLogout() {
// 	// create test jwt
// }

// func (lt *LoginSuite) TestValidate() {

// }

// func (lt *LoginSuite) TestRegister() {

// }

// func (lt *LoginSuite) TestCreateUser() {

// }

// func (lt *LoginSuite) TestModifyUser() {

// }

// func (lt *LoginSuite) TestDeleteUser() {

// }
