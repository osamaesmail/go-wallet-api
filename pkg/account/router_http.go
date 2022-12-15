package account

import (
	"github.com/gin-gonic/gin"
)

func NewHTTPRouter(apiGroup *gin.RouterGroup, controller HTTPTransport) {
	txAPIGroup := apiGroup.Group("/accounts")
	txAPIGroup.POST("create", gin.WrapH(controller.Create()))
	txAPIGroup.POST("list", gin.WrapH(controller.List()))
}
