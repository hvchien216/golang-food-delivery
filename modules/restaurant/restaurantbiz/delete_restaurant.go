package restaurantbiz

import (
	"context"
	"errors"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type DeleteRestaurantStore interface {
	FindRestaurantById(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(
	ctx context.Context,
	id int,
) error {

	result, err := biz.store.FindRestaurantById(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if result.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return err
	}

	return nil
}
