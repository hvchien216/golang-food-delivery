package ginrestaurantlike

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurantlikebiz "food_delivery/modules/restaurantlike/biz"
	restaurantlikemodel "food_delivery/modules/restaurantlike/model"
	restaurantlikestorage "food_delivery/modules/restaurantlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// POST /v1/restaurants/:id/like

func LikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurantlikemodel.Like{
			RestaurantId: int(uid.GetLocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		//increaseLikeCountStore := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		//biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, increaseLikeCountStore, appCtx.GetPubsub())
		biz := restaurantlikebiz.NewUserLikeRestaurantBiz(store, appCtx.GetPubsub())

		if err = biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)

		}

		c.JSON(http.StatusOK, common.SimpleSucessResponse(true))
	}
}
