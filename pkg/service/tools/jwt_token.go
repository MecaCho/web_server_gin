package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
	"math"
	"strconv"
	"time"

	"web_server_gin/pkg/model/authmgr"
)

// GetUserToken ...
func GetUserToken(user authmgr.User, expireMinutes time.Duration) (token string, timeExpire time.Time, err error) {
	timeExpire = time.Now().Add(expireMinutes)
	expiredTime := timeExpire.Unix()

	rolesList := make([]string, 0)
	for k := range user.Roles {
		rolesList = append(rolesList, user.Roles[k].Name)
	}
	if len(rolesList) == 0 {
		rolesList = append(rolesList, string(authmgr.RoleGuest))
	}

	groupList := make([]string, 0)
	for k := range user.Groups {
		groupList = append(groupList, strconv.Itoa(int(user.Groups[k].ID)))
	}
	tokenNew := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_name":  user.Name,
		"expired_at": expiredTime,
		"roles":      rolesList,
		"groups":     groupList,
	})

	token, _ = tokenNew.SignedString([]byte("secret"))

	return
}

// AuthValidate ...
func AuthValidate(tokenStr string) (err error) {
	if tokenStr == "" {
		return fmt.Errorf("not authorization")
	}

	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("not authorization")
		}
		return []byte("secret"), nil
	})
	if !token.Valid {
		return fmt.Errorf("not authorization, token invalidated")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userName := claims["user_name"].(string)
	fmt.Printf("parase result : %+v, token user name : %s", ok, userName)

	sec, dec := math.Modf(claims["expired_at"].(float64))
	glog.Infoln("toke paraser time : ", time.Unix(int64(sec), int64(dec*(1e9))))

	glog.Infof("token parase claims : %+v.", claims)
	glog.Infoln("token header: ", token.Header, token.Method, token.Raw, token.Signature)
	return
}
