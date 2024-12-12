package api

import (
	v1 "github.com/DnsUnlock/Dpanel/backend/api/v1"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	v1.Router(r.Group("/v1"))
}
