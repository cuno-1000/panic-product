package middleware

import (
	"errors"
	"github.com/asim/go-micro/v3/logger"
	"github.com/cuno-1000/panic-product/api-gateway/token/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"time"
)

const (
	TokenType    = "bearer"
	AppGuardName = "app"
	Secret       = "nothememeetingbutteambuilding"

	JwtBlacklistGracePeriod = 10
	RefreshGracePeriod      = 1800
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusBadRequest, "Cannot read token")
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(services.TokenType)+1:]

		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(Secret), nil
		})
		if err != nil || services.JwtService.IsInBlacklist(tokenStr) {
			c.JSON(http.StatusBadRequest, "Invalid Token")
			c.Abort()
			return
		}

		claims := token.Claims.(*services.CustomClaims)
		if claims.Issuer != GuardName {
			c.JSON(http.StatusBadRequest, "Cannot read token")
			c.Abort()
			return
		}

		// token 续签
		if claims.ExpiresAt-time.Now().Unix() < RefreshGracePeriod {
			lock := services.Lock("refresh_token_lock", JwtBlacklistGracePeriod)
			if lock.Get() {
				userId, err := services.JwtService.GetUserId(GuardName, claims.Id)
				if err != nil {
					logger.Fatal(err)
					//global.App.Log.Error(err.Error())
					lock.Release()
				} else {
					tokenData, _, _ := services.JwtService.CreateToken(GuardName, userId)
					c.Header("new-token", tokenData.AccessToken)
					c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
					_ = services.JwtService.JoinBlackList(token)
				}
			}
		}

		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}

func IdFromJWT(c *gin.Context) (id uint64, err error) {
	tokenStr := c.Request.Header.Get("Authorization")
	if tokenStr == "" {
		c.JSON(http.StatusBadRequest, "Cannot read token")
		c.Abort()
		return 0, errors.New("cannot read token")
	}
	tokenStr = tokenStr[len(services.TokenType)+1:]

	token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil || services.JwtService.IsInBlacklist(tokenStr) {
		c.JSON(http.StatusBadRequest, "Cannot read token")
		c.Abort()
		return 0, err
	}
	return strconv.ParseUint(token.Claims.(*services.CustomClaims).Id, 10, 32)
}
