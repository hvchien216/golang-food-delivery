package restaurantlikebiz

import (
	"context"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/modules/restaurantlike/model"
)

type ListUsersLikeRestaurantStore interface {
	GetUsersLikeRestaurant(ctx context.Context,
		condition map[string]interface{},
		filter *restaurantlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUsersLikeRestaurantBiz struct {
	store ListUsersLikeRestaurantStore
}

func NewListUsersLikeRestaurantBiz(store ListUsersLikeRestaurantStore) *listUsersLikeRestaurantBiz {
	return &listUsersLikeRestaurantBiz{store: store}
}

func (biz *listUsersLikeRestaurantBiz) ListUsers(ctx context.Context,
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {

	users, err := biz.store.GetUsersLikeRestaurant(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantlikemodel.EntityName, err)
	}

	return users, nil
}
