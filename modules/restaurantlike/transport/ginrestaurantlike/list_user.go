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

func ListUsersLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		filter := restaurantlikemodel.Filter{
			RestaurantId: int(uid.GetLocalID()),
		}

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{"oke": 1})
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{"oke": 1})
			return
		}

		paging.Fulfill()

		//myArr := []string{}
		//
		//fmt.Println(myArr[0])

		store := restaurantlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantlikebiz.NewListUsersLikeRestaurantBiz(store)

		users, err := biz.ListUsers(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		for i := range users {
			users[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(users, paging, filter))
	}
}
