package middleware

import (
	"os"
	"strconv"
	"time"

	"github.com/Fishmansky/noteflow/inits"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/twinj/uuid"
)

type TokenDetails struct {
	Token string
	Uuid  string
	Exp   int64
}

func Authenticate(c *gin.Context) bool {
	return false
}

func CreateAccessToken(userID uint) (*TokenDetails, error) {
	accessToken := &TokenDetails{}
	accessToken.Uuid = uuid.NewV4().String()
	accessToken.Exp = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized":  "true",
		"access_uuid": accessToken.Uuid,
		"user_id":     userID,
		"exp":         accessToken.Exp,
	})
	var err error
	accessToken.Token, err = token.SignedString([]byte(os.Getenv("SECRET")))
	return accessToken, err
}

func CreateRefreshToken(userID uint) (*TokenDetails, error) {
	refreshToken := &TokenDetails{}
	refreshToken.Uuid = uuid.NewV4().String()
	refreshToken.Exp = time.Now().Add(time.Hour * 24 * 90).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized":   "true",
		"refresh_uuid": refreshToken.Uuid,
		"user_id":      userID,
		"exp":          refreshToken.Exp,
	})
	var err error
	refreshToken.Token, err = token.SignedString([]byte(os.Getenv("SECRET")))
	return refreshToken, err
}

func SaveToken(userID uint, t *TokenDetails) error {
	Expires := time.Unix(t.Exp, 0)
	now := time.Now()

	err := inits.Redis.Set(t.Uuid, strconv.Itoa(int(userID)), Expires.Sub(now)).Err()
	if err != nil {
		return err
	}
	return nil
}

func RefreshTokens(c *gin.Context) {

}

func DeleteToken(c *gin.Context) {

}
