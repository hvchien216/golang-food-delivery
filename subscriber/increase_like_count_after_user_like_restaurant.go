package subscriber

import (
	"context"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/modules/restaurant/restaurantstorage"
)

type HasRestaurantId interface {
	GetRestaurantId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubsub().Subscribe(ctx, common.TopicUserLikeRestaurant)
	store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
	go func() {
		for {
			msg := <-c
			likeData := (msg.Data()).(HasRestaurantId)
			store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}
