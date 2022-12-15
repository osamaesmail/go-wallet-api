package transaction

import (
	"github.com/gin-gonic/gin"
)

func NewHttpRouter(apiGroup *gin.RouterGroup, controller HTTPTransport) {
	txAPIGroup := apiGroup.Group("/transactions")
	txAPIGroup.POST("create", gin.WrapH(controller.Create()))
	txAPIGroup.POST("list", gin.WrapH(controller.List()))
}
