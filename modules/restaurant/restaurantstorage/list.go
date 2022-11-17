package restaurantstorage

import (
	"context"
	"food_delivery/common"
	"food_delivery/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	condition map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant
	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where(condition).
		Where("status in (1)")

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city = ?", v.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit
	//log.Fatalln("paigngggggg", paging)
	if err := db.
		Offset(int(offset)).
		Limit(int(paging.Limit)).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
