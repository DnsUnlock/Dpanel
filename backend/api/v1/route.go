package v1

import (
	"github.com/DnsUnlock/Dpanel/backend/api/v1/captcha"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	captcha.Router(r.Group("/captcha"))
}
