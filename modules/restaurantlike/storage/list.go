package restaurantlikestorage

import (
	"context"
	"fmt"
	"food_delivery/common"
	restaurantlikemodel "food_delivery/modules/restaurantlike/model"
	"github.com/btcsuite/btcutil/base58"
	"time"
)

func (s *sqlStore) GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id"`
		LikedCount   int `gorm:"column:count"`
	}

	var listLike []sqlData

	if err := s.db.Table(restaurantlikemodel.Like{}.TableName()).Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").
		Find(&listLike).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range listLike {
		result[item.RestaurantId] = item.LikedCount
	}

	return result, nil
}

var timeLayout = "2006-01-02T15:04:05.99999"
var timeLayoutNoSecond = "2006-01-02T15:04:05"

func (s *sqlStore) GetUsersLikeRestaurant(ctx context.Context,
	condition map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]common.SimpleUser, error) {
	var result []restaurantlikemodel.Like
	db := s.db.Table(restaurantlikemodel.Like{}.TableName()).Where(condition)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db = db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	//for i := range moreKeys {
	//	db = db.Preload(moreKeys[i])
	//}

	db = db.Preload("User")

	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		// format with timeLayoutNoSecond: becuz SQL just understand this format for time "2006-01-02T15:04:05"
		db = db.Where("created_at < ?", timeCreated.Format(timeLayoutNoSecond))
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(int(offset))

	}

	if err := db.
		Limit(int(paging.Limit)).
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	users := make([]common.SimpleUser, len(result))

	for i, item := range result {
		result[i].User.CreatedAt = item.CreatedAt
		result[i].User.UpdatedAt = nil
		users[i] = *result[i].User // when len of two arr not equal (mean: user may is null) => system crashing
		if i == len(result)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return users, nil
}
