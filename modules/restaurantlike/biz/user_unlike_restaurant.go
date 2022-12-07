package restaurantlikebiz

import (
	"context"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/modules/restaurantlike/model"
)

type UserUnLikeRestaurantStore interface {
	Delete(ctx context.Context, userId, restautantId int) error
	Find(ctx context.Context, conditions map[string]interface{}) (*restaurantlikemodel.Like, error)
}

type DecreaseLikeCountStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnLikeRestaurantBiz struct {
	store                  UserUnLikeRestaurantStore
	decreaseLikeCountStore DecreaseLikeCountStore
}

func NewUserUnLikeRestaurantBiz(store UserUnLikeRestaurantStore, decreaseLikeCountStore DecreaseLikeCountStore) *userUnLikeRestaurantBiz {
	return &userUnLikeRestaurantBiz{store: store, decreaseLikeCountStore: decreaseLikeCountStore}
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

	go func() {
		defer common.AppRecover()
		// side effect
		_ = biz.decreaseLikeCountStore.DecreaseLikeCount(ctx, restaurantId)
	}()

	return nil
}
