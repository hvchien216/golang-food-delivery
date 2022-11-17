package restaurantstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindRestaurantById(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		// case: error from DB
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
