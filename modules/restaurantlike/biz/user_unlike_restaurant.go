package restaurantlikebiz

import (
	"context"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/modules/restaurantlike/model"
	"food_delivery/pubsub"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restautantId int) error
	Find(ctx context.Context, conditions map[string]interface{}) (*restaurantlikemodel.Like, error)
}

//type DecreaseLikeCountStore interface {
//	DecreaseLikeCount(ctx context.Context, id int) error
//}

type userUnLikeRestaurantBiz struct {
	store UserUnLikeRestaurantStore
	//decreaseLikeCountStore DecreaseLikeCountStore,
	pubsub pubsub.Pubsub
}

func NewUserUnLikeRestaurantBiz(
	store UserUnLikeRestaurantStore,
	//decreaseLikeCountStore DecreaseLikeCountStore
	pubsub pubsub.Pubsub) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{
		store: store,
		//decreaseLikeCountStore: decreaseLikeCountStore
		pubsub: pubsub,
	}
}

func (biz *userUnLikeRestaurantBiz) UnLikeRestaurant(
	ctx context.Context,
	userId, restaurantId int,
) error {

	_, err := biz.store.Find(ctx, map[string]interface{}{"user_id": userId, "restaurant_id": restaurantId})

	if err != nil {
		return restaurantlikemodel.ErrAlreadyUnLikedRestaurant()
	}

	err = biz.store.Delete(ctx, userId, restaurantId)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// Side Effect
	//without Job
	//go func() {
	//
	//	defer common.AppRecover()
	//	_ = biz.decreaseLikeCountStore.DecreaseLikeCount(ctx, restaurantId)
	//}()

	//with Job
	//go func() {
	//	defer common.AppRecover()
	//	job := asyncjob.NewJob(func(ctx context.Context) error {
	//		return biz.decreaseLikeCountStore.DecreaseLikeCount(ctx, restaurantId)
	//
	//	})
	//	_ = asyncjob.NewGroup(true, job).Run(ctx)
	//}()

	// Pubsub
	biz.pubsub.Publish(
		ctx,
		common.TopicUserDislikeRestaurant,
		pubsub.NewMessage(&restaurantlikemodel.Like{
			RestaurantId: restaurantId,
			UserId:       userId,
		}))

	return nil
}
