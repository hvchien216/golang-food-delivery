package restaurantlikebiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/component/asyncjob"
	restaurantlikemodel "food_delivery/modules/restaurantlike/model"
	"time"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurantlikemodel.Like) error
	Find(ctx context.Context, conditions map[string]interface{}) (*restaurantlikemodel.Like, error)
}

type IncreaseLikeCountStore interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeRestaurantBiz struct {
	store             UserLikeRestaurantStore
	increaseLikeStore IncreaseLikeCountStore
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, increaseLikeStore IncreaseLikeCountStore) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, increaseLikeStore: increaseLikeStore}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(
	ctx context.Context,
	data *restaurantlikemodel.Like,
) error {

	_, err := biz.store.Find(ctx, map[string]interface{}{"user_id": data.UserId, "restaurant_id": data.RestaurantId})

	if err == nil {
		return restaurantlikemodel.ErrAlreadyLikedRestaurant()
	}

	err = biz.store.Create(ctx, data)

	if err != nil {
		return restaurantlikemodel.ErrCannotLikeRestaurant(err)
	}

	// Side Effect

	// without Job
	//go func() {
	//	defer common.AppRecover()
	//	_ = biz.increaseLikeStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//}()

	// with Job
	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.increaseLikeStore.IncreaseLikeCount(ctx, data.RestaurantId)

		})

		job.SetRetryDurations([]time.Duration{time.Second * 3})

		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()

	return nil
}
