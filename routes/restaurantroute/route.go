package restaurantroute

import (
	"food_delivery/component/appctx"
	"food_delivery/middleware"
	"food_delivery/modules/restaurant/restauranttransport/ginrestaurant"
	"food_delivery/modules/restaurantlike/transport/ginrestaurantlike"
	"github.com/gin-gonic/gin"
)

func Routes(v1 *gin.RouterGroup, appCtx appctx.AppContext) {
	restaurants := v1.Group("/restaurants", middleware.RequireAuth(appCtx))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUsersLikeRestaurant(appCtx))
		restaurants.POST("/:id/like", ginrestaurantlike.LikeRestaurant(appCtx))
		restaurants.DELETE("/:id/unlike", ginrestaurantlike.UnLikeRestaurant(appCtx))
	}
}
