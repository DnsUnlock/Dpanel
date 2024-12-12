package captcha

import (
	"github.com/DnsUnlock/Dpanel/backend/model/response"
	"github.com/DnsUnlock/Dpanel/backend/utils/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckCaptcha(c *gin.Context) {
	var data struct {
		Point string `json:"point"`
		Key   string `json:"key"`
	}
	if err := c.ShouldBindBodyWithJSON(&data); err != nil {
		c.JSON(http.StatusOK, response.Error(err.Error(), nil))
		return
	}
	checkCaptcha, err := captcha.CheckCaptcha(data.Point, cache[data.Key])
	if err != nil {
		c.JSON(http.StatusOK, response.Error(err.Error(), nil))
	}
	if checkCaptcha {
		c.JSON(http.StatusOK, response.Success("验证成功", gin.H{
			"token": "token",
		}))
	} else {
		c.JSON(http.StatusOK, response.Error("验证失败", nil))
	}
}
