package jwe

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/DnsUnlock/Dpanel/backend/utils/jwe"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

// CheckJwt 通过 Jwt 获取信息
func CheckJwt(c *gin.Context) {

	// 从请求头内获取 Jwt
	authorization := c.GetHeader("Authorization")
	authorization = strings.Replace(authorization, "Bearer ", "", -1)
	session := c.GetHeader("Session")
	token, err := jwe.GetToken(authorization)
	if err != nil {
		c.JSON(401, gin.H{
			"code": "S00",
			"msg":  "无效的 Jwt",
		})
		c.Abort()
		return
	}
	if !verifyHash(token.Finger, session) {
		c.JSON(401, gin.H{
			"code": "S00",
			"msg":  "无效的 Session",
		})
		c.Abort()
		return
	}
	c.Next()
}

func verifyHash(finger, hash string) bool {
	// 获取当前 Unix 时间戳
	currentTime := time.Now().Unix()
	maxTimeDifference := int64(5) // 最大时间偏差为 5 秒，前后共计 10 秒

	// 创建一个哈希比较函数
	checkHash := func(offset int64) bool {
		timestamp := currentTime + offset
		dataToHash := fmt.Sprintf("%s%d", finger, timestamp)

		// 计算 MD5 哈希值
		hashBytes := md5.Sum([]byte(dataToHash))
		calculatedHash := hex.EncodeToString(hashBytes[:])
		fmt.Println(calculatedHash)
		// 比较哈希值
		return strings.EqualFold(calculatedHash, hash)
	}

	// 从当前时间戳向两边扩展
	for i := int64(0); i <= maxTimeDifference; i++ {
		if checkHash(i) || checkHash(-i) {
			return true
		}
	}

	return false
}
