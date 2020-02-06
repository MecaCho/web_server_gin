package util

import (
	"fmt"
	"github.com/golang/glog"
	"math/rand"
	"strconv"
	"time"
)

// GenPwd generate a random password
func GenPwd(num int) string {
	rand.NewSource(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	res := make([]byte, num)
	seedStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < num; i++ {
		seq := rand.Intn(len(seedStr))
		res[i] = seedStr[seq]
	}
	return string(res)
}

// TimeAfterMinutes check time after
func TimeAfterMinutes(timeObj time.Time, mins int) bool {
	timeNow := time.Now()
	minsStr := fmt.Sprintf("%sm", strconv.Itoa(mins))
	minDuration, _ := time.ParseDuration(minsStr)
	timDelay := timeObj.Add(minDuration)
	ret := timeNow.After(timDelay)
	if ret {
		glog.Infof("Time after : %+v", ret)
	}
	return ret
}
