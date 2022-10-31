package services

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/asim/go-micro/v3/logger"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

func init() {
	RedisDb = redisInit()
}

type RedisConnect struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

func redisInit() *redis.Client {
	var redisInfo = &RedisConnect{
		Host:     "127.0.0.1",
		Port:     6379,
		DB:       1,
		Password: "",
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisInfo.Host + ":" + strconv.Itoa(redisInfo.Port),
		Password: redisInfo.Password, // no password set
		DB:       redisInfo.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Fatal(err)
		return nil
	}
	return client
}

var RedisDb *redis.Client

type Response struct {
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

type jwtService struct {
}

var JwtService = new(jwtService)

// CustomClaims 自定义 Claims
type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType    = "bearer"
	AppGuardName = "app"

	Ttl                     = 43200
	JwtBlacklistGracePeriod = 10
	Secret                  = "nothememeetingbutteambuilding"
)

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (jwtService *jwtService) CreateToken(GuardName string, userId uint64) (tokenData TokenOutPut, err error, token *jwt.Token) {
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + Ttl,
				Id:        strconv.FormatUint(userId, 10),
				Issuer:    GuardName,
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(Secret))

	tokenData = TokenOutPut{
		tokenStr,
		int(Ttl),
		TokenType,
	}
	return
}

func MD5(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

func (jwtService *jwtService) getBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + MD5([]byte(tokenStr))
}

func (jwtService *jwtService) JoinBlackList(token *jwt.Token) (err error) {
	nowUnix := time.Now().Unix()
	//timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt - nowUnix) * time.Second
	a := token.Claims.(*CustomClaims).ExpiresAt - nowUnix
	timer := time.Duration(a) * time.Second
	err = RedisDb.SetNX(context.Background(), jwtService.getBlackListKey(token.Raw), nowUnix, timer).Err()
	return
}

func (jwtService *jwtService) IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := RedisDb.Get(context.Background(), jwtService.getBlackListKey(tokenStr)).Result()
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	if time.Now().Unix()-joinUnix < JwtBlacklistGracePeriod {
		return false
	}
	return true
}

func (jwtService *jwtService) GetUserId(GuardName string, id string) (userId uint64, err error) {
	switch GuardName {
	case AppGuardName:
		return strconv.ParseUint(id, 10, 64)
	default:
		err = errors.New("guard " + GuardName + " does not exist")
	}
	return
}
