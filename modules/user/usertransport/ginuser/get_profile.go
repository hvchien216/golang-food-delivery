package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data))
	}
}
