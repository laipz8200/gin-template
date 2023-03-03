package middlewares

import (
	"_template_/api/schemas"
	"_template_/config"
	"_template_/constants"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// A struct to store the user info from the jwt
type UserInfo struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}

// A function to extract the jwt from the "X-Authorization" header
func getJWT(c *gin.Context) (string, error) {
	header := c.GetHeader("X-Authorization")
	if header == "" {
		return "", fmt.Errorf("no authorization header")
	}
	return header, nil
}

// A function to parse the jwt and get the user info from it
func getUserInfo(jwtString string) (*UserInfo, error) {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret()), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	userInfo := &UserInfo{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userInfo.ID = claims["id"].(string)
		userInfo.Username = claims["username"].(string)
		roles := claims["roles"].([]any)
		userInfo.Roles = make([]string, len(roles))
		for i, v := range roles {
			userInfo.Roles[i] = v.(string)
		}
	}
	return userInfo, nil
}

// A middleware function to authenticate the request and set the user info in the context
func AuthMiddleware(c *gin.Context) {
	jwtString, err := getJWT(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, schemas.ErrMsg{
			Code:  http.StatusUnauthorized,
			Error: err.Error(),
		})
		return
	}
	userInfo, err := getUserInfo(jwtString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, schemas.ErrMsg{
			Code:  http.StatusUnauthorized,
			Error: err.Error(),
		})
		return
	}
	c.Set(constants.KEY_CURRENT_USER, userInfo)
	c.Next()
}
