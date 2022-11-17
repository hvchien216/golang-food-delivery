package middleware

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"github.com/gin-gonic/gin"
)

func Recover(ac appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// avoid case that response result type is text
				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// re-enable panicking mechanism for `Gin lib` cuz `Gin` has its own recovery
					panic(err)
					return
				}

				// `err.(error)` just return a error cuz `err` is of type result `recover()`
				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}
		}()

		c.Next()
	}
}
