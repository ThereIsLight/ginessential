package common

import (
	"ginEssential/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_crect")

// jwt正文部分
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 登录成功之后使用该方法发放Token
func ReleaseToken(user model.User) (string, error) {
	expirationtime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims :jwt.StandardClaims{
			ExpiresAt: expirationtime.Unix(), // 过期时间
			IssuedAt: time.Now().Unix(), // 发放时间
			Issuer: "YG", // 发放者
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)  // 加密方式已经加密内容
	tokenString, err := token.SignedString(jwtKey)  // 使用密钥生成token
	if err != nil {
		return "", err
	}
	return tokenString, nil
	/*eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
	eyJVc2VySWQiOjQsImV4cCI6MTU5NTQwOTI0OSwiaWF0IjoxNTk0ODA0NDQ5LCJpc3MiOiJZRyIsInN1YiI6InVzZXIgdG9rZW4ifQ.
	GA2e2jIcCCK6trVy415Gu0jK0nt1o8ES1Ym8ZOT5Cs8*/
}