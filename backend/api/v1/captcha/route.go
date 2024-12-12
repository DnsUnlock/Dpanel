package captcha

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	// 验证码相关
	r.GET("", GetCaptcha)    // 获取验证码
	r.POST("", CheckCaptcha) // 验证验证码
}
