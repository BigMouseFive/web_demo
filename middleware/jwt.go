package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/web_demo/v2/log"
	"github.com/web_demo/v2/model/common"
	"github.com/web_demo/v2/model/demo"
	demo2 "github.com/web_demo/v2/service/demo"
	"net/http"
	"time"
)

var (
	JwtMW *jwt.GinJWTMiddleware
)
var identityKey = "uuTTKkInnJKhNN"

// CreateJWT 创建jwt
func CreateJWT() {
	var err error
	JwtMW, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*common.JWTPayload); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &common.JWTPayload{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginInfo demo.LoginRequest
			if err := c.ShouldBind(&loginInfo); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			status, _ := demo2.LoginSvc.Login(loginInfo)
			if status == http.StatusOK {
				return &common.JWTPayload{
					UserName: loginInfo.Account,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// 后续权限相关的可以放在此处
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, common.Response{Msg: message})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, common.Response{Result: gin.H{
				"token":  token,
				"expire": expire.Format(time.RFC3339)},
			})
		},
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(code, common.Response{})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, common.Response{Result: gin.H{
				"token":  token,
				"expire": expire.Format(time.RFC3339)},
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Sugar.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := JwtMW.MiddlewareInit()

	if errInit != nil {
		log.Sugar.Fatal("JWTMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
}
