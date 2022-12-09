package subscriber

import (
	"context"
	"food_delivery/component/appctx"
	"food_delivery/modules/restaurant/restaurantstorage"
	"food_delivery/pubsub"
)

func RunDecreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease like count after user unlikes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

			likeData := (message.Data()).(HasRestaurantId)
			return store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
