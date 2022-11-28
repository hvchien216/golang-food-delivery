package ginrestaurant

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/modules/restaurant/restaurantbiz"
	"food_delivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		go func() {
			defer common.AppRecover()
			panic("aaaa")
		}()

		uid, err := common.FromBase58(c.Param("id"))

		// This is an error from Go standard lib, so we need to wrap it by common.ErrInvalidRequest
		// cuz this error is not normalized
		if err != nil {
			// NOTICE:
			// we should just set `panic` in the transportation/controller layer
			// If we set `panic` in the business/services layer, because `panic`'s mechanism
			// will stop any code below it, so we might miss some logic in it
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		result, err := biz.GetRestaurant(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			// Any err thrown from Biz belongs to Application error
			panic(err)
		}

		result.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSucessResponse(result))
	}
}
