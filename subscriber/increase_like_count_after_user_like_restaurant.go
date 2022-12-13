package subscriber

import (
	"context"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/modules/restaurant/restaurantstorage"
	"food_delivery/pubsub"
	"food_delivery/skio"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	GetUserId() int
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

func RunIncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase like count after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

			likeData := (message.Data()).(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func EmitRealtimeAfterUserLikeRestaurant(appCtx appctx.AppContext, rtEngine skio.RealtimeEngine) consumerJob {
	return consumerJob{
		Title: "Emit realtime after user likes restaurant",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			likeData := (message.Data()).(HasRestaurantId)
			store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())

			result, _ := store.FindRestaurantById(context.Background(), map[string]interface{}{"id": likeData.GetRestaurantId()})

			ownerRestaurantId := result.GetOwnerId()

			//emit to Restaurant's owner
			return rtEngine.EmitToUser(ownerRestaurantId, string(message.Channel()), message.Data())
		},
	}
}
