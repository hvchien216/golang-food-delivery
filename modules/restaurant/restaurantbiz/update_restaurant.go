package restaurantbiz

import (
	"context"
	"errors"
	"food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type UpdateRestaurantStore interface {
	FindRestaurantById(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *restaurantmodel.RestaurantUpdate,
	) error
}

type updateRestaurantBiz struct {
	store     UpdateRestaurantStore
	requester common.Requester
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore, requester common.Requester) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store, requester: requester}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(
	ctx context.Context,
	id int,
	data *restaurantmodel.RestaurantUpdate,
) error {

	result, err := biz.store.FindRestaurantById(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if result.Status == 0 {
		return errors.New("data deleted")
	}

	// only admin & restaurant's owner can update data
	if biz.requester.GetRole() != "admin" && result.UserId != biz.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
