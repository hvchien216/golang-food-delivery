package restaurantrepo

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

type ListRestaurantStore interface {
	ListDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]restaurantmodel.Restaurant, error)
}

// to optimize algorithm: []restaurantlikemodel.Like => map[int]int
// becuz when Join restaurant_likes table & restaurant = O(n^2)
// when use Map, we can reduce the algorithm = O(n)
type RestaurantLikeStore interface {
	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
}

type listRestaurantRepo struct {
	store     ListRestaurantStore
	likeStore RestaurantLikeStore
}

func NewListRestaurantRepo(store ListRestaurantStore,
	likeStore RestaurantLikeStore) *listRestaurantRepo {
	return &listRestaurantRepo{
		store:     store,
		likeStore: likeStore,
	}
}

func (repo *listRestaurantRepo) ListRestaurant(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {

	result, err := repo.store.ListDataByCondition(ctx, nil, filter, paging, "User")

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	//ids := make([]int, len(result))
	//
	//for i := range result {
	//	ids[i] = result[i].Id
	//}
	//
	//mapLikesResponse, err := repo.likeStore.GetRestaurantLikes(ctx, ids)
	//
	//// Tại dây: vì nếu apply theo cách cũ (JOIN) thì nếu bảng được JOIN lỗi => chết luôn
	//// Còn ở đây: nếu mà có err thì chỉ nên show ra msg (Ex: không thể lấy lượt like) không nên để nó bị ảnh hưởng nếu bị lỗi
	//// ====> we just log
	//// Rất thích hợp khi build Mircoservices
	//if err != nil {
	//	//return nil, common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	//	log.Println("Cannot get restaurant likes", err)
	//}
	//
	//if v := mapLikesResponse; v != nil {
	//	for i, item := range result {
	//		result[i].LikedCount = mapLikesResponse[item.Id]
	//	}
	//}

	return result, nil
}
