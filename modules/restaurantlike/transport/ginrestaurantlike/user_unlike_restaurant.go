package ginrestaurantlike

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantlikebiz "food_delivery/modules/restaurantlike/biz"
	restaurantlikestorage "food_delivery/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /v1/restaurants/:id/unlike

func UnLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		//decreaseLikeCountRestaurant := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		pubsub := appCtx.GetPubsub()
		//biz := restaurantlikebiz.NewUserUnLikeRestaurantBiz(store, decreaseLikeCountRestaurant)
		biz := restaurantlikebiz.NewUserUnLikeRestaurantBiz(store, pubsub)

		if err = biz.UnLikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
