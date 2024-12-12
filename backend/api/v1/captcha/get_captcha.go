package captcha

import (
	"github.com/DnsUnlock/Dpanel/backend/model/response"
	"github.com/DnsUnlock/Dpanel/backend/utils/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
)

var cache = make(map[string][]byte)

func GetCaptcha(c *gin.Context) {
	captchaData, checkData, err := captcha.GenerateCaptcha()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "生成验证码失败",
		})
		return
	}
	cache[captchaData.CaptchaKey] = checkData.CaptchaByte
	c.JSON(http.StatusOK, response.Success(
		"生成验证码成功",
		captchaData,
	))
}
