package restaurantbiz

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type FindSingleRestaurantStore interface {
	FindRestaurantById(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBiz struct {
	store FindSingleRestaurantStore
}

func NewGetRestaurantBiz(store FindSingleRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(
	ctx context.Context,
	id int) (*restaurantmodel.Restaurant, error) {

	result, err := biz.store.FindRestaurantById(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			//return nil, common.ErrEntityNotFound(restaurantmodel.EntityName, err)
			return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)

		}
		// OR: able to throw err `sth went wrong with server`
		return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
	}

	// for case soft deleted (mean: can't retrieve record when status == 0)
	if result.Status == 0 {
		// FOR CASE Security:
		//return nil, common.ErrCannotGetEntity(restaurantmodel.EntityName, err)
		return nil, common.ErrEntityDeleted(restaurantmodel.EntityName, err)
	}
	//in b
	//a
	// enterring: s
	// leaving: s
	//
	return result, err
}
