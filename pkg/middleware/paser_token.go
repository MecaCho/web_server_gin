package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"math"
	"strings"
	"time"
	"web_server_gin/pkg/util"
)

// UserInfo user info
type UserInfo struct {
	Name   string   `json:"name"`
	Groups []string `json:"groups"`
	Roles  []string `json:"roles"`
}

// CheckToken check authorization
func ParserToken(tokenStr string) (user UserInfo, err error) {
	if strings.Contains(tokenStr, "Bearer") {
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer")
	}
	tokenStr = strings.TrimSpace(tokenStr)
	if tokenStr == "" {
		err = fmt.Errorf("User unauthorized, authorization token not found.")
		return
	}
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		return []byte("secret"), nil
	})
	if token == nil || !token.Valid {
		err = fmt.Errorf("User unauthorized, token invalid.")
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userInf := claims["user_name"]
	expireInf := claims["expired_at"]
	if !ok || expireInf == nil {
		err = fmt.Errorf("User unauthorized, token claims invalid.")
		return
	}
	user.Roles = ParseCalims(claims, "roles")
	user.Groups = ParseCalims(claims, "groups")

	sec, dec := math.Modf(claims["expired_at"].(float64))
	expiredTime := time.Unix(int64(sec), int64(dec*(1e9)))
	glog.Infof("Token expired time : %v, user name: %v, roles: %+v, groups: %+v.", expiredTime, userInf, user.Roles, user.Groups)
	if util.TimeAfterMinutes(expiredTime, 0) {
		err = fmt.Errorf("User unauthorized, token has expired.")
		return
	}
	user.Name = userInf.(string)
	return
}

// ParseCalims parse token calims
func ParseCalims(claims jwt.MapClaims, key string) (roles []string) {
	roleInf := claims[key]
	if roleInf != nil {
		rolesInterf := roleInf.([]interface{})
		for k := range rolesInterf {
			roles = append(roles, rolesInterf[k].(string))
		}
	}
	return
}
