package ginrestaurant

import (
	"food_delivery/component/appctx"
	"food_delivery/modules/restaurant/restaurantbiz"
	"food_delivery/modules/restaurant/restaurantmodel"
	"food_delivery/modules/restaurant/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"ok": 1})
	}
}
