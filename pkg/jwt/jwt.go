package jwt

import (
	"errors"
	"forum/common"
	fError "forum/utils/forum_error"
	"time"

	"go.uber.org/zap"

	"github.com/dgrijalva/jwt-go"
)

const aTokenExpireDuration = time.Hour * 24830
const rTokenExpireDuration = time.Hour * 24 * 30

var secret = []byte("Kassadin")

type Claim struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成token
func GenToken(userId int64, username string) (aToken, rToken string, err error) {
	c := Claim{UserId: userId, Username: username, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(aTokenExpireDuration).Unix(), Issuer: common.APPName}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	aToken, err = token.SignedString(secret)
	if err != nil {
		zap.L().Error("Generate access token failed.", zap.Error(err))
		return aToken, rToken, fError.New(fError.CodeSystemError)
	}
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{ExpiresAt: time.Now().Add(rTokenExpireDuration).Unix(), Issuer: common.APPName}).SignedString(secret)
	if err != nil {
		zap.L().Error("Generate refresh token failed.", zap.Error(err))
		return aToken, rToken, fError.New(fError.CodeSystemError)
	}
	return aToken, rToken, err
}

// keyFunc 生成token时使用
func keyFunc(token *jwt.Token) (i interface{}, err error) {
	return secret, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*Claim, error) {
	// 解析token到自定义结构体中，第三个参数是处理盐值到函数
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, keyFunc)
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	// 校验token
	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims, nil
	}
	return nil, fError.New(fError.CodeUserNotLogin)
}

// RefreshToken 刷新Token
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	// 解析token，无效直接返回
	if _, err = jwt.Parse(rToken, keyFunc); err != nil {
		zap.L().Error("Parse token failed.", zap.Error(err))
		return "", "", fError.New(fError.CodeSystemError)
	}
	// 从旧的access token中解析claims数据，用于生成新token
	var claim Claim
	// 解析旧的access token，如果返回超期的错误，则生成新的token
	_, err = jwt.ParseWithClaims(aToken, claim, keyFunc)
	var vErr *jwt.ValidationError
	_ = errors.As(err, &vErr)
	if vErr.Errors == jwt.ValidationErrorExpired {
		zap.L().Info("Refresh token success.", zap.String("Username", claim.Username))
		return GenToken(claim.UserId, claim.Username)
	} else {
		// 其他错误，返回请重新登陆
		zap.L().Error("Refresh token failed.", zap.Error(err))
		return "", "", fError.New(fError.CodeUserNotLogin)
	}
}
