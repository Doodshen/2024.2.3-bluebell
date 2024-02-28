package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 2 定义加盐 定义一个用于签名的字符串：
var mySercet = []byte("kingshen")

// 4 生成token有效时间
const TokenExpireDuration = time.Hour * 2

// 1 定义自己的Claim
type MyClaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// Valid implements jwt.Claims.

// 3 生成jwt
// GenToken生成Jwt
func GenToken(userid int64, username string) (string, error) {
	//创建一个自己的声明Claim
	c := MyClaims{
		userid,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "bluebell",                                 //签发人
		},
	}

	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	//使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(mySercet)
}

// 5 解析token
// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法----不仅校验外还直接写到我们的变量当中去
	var mc = new(MyClaims)

	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return mySercet, nil
	})
	if err != nil {
		return nil, err
	}

	// 对token对象中的Claim进行类型断言--转换为变量
	if token.Valid { // 校验token
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
