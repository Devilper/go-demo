package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-demo/global"
	"net/http"
	"time"
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// SignKey 常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenInvalid     error  = errors.New("Couldn't handle this token")
	SignKey          string = "devil"
)

// CustomClaims 载荷, 可以加一些自己需要的信息
type CustomClaims struct {
	ID    string `json:"userId"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// GetSignKey 获取SignKey
func GetSignKey() string {
	return SignKey
}

// SetSignKey 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// Auth 中间件，检测token
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := global.NewResult(c)
		token := c.Request.Header.Get("token")
		if token == "" {
			result.Error(http.StatusBadRequest, "请求未携带token,无权访问")
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				result.Error(http.StatusBadRequest, "授权已过期")
				c.Abort()
				return
			}
			result.Error(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

// ParseToken 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	} else {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, TokenInvalid
}

// CreatToken 生成一个token
func (j *JWT) CreatToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreatToken(*claims)
	}
	return "", TokenInvalid
}
