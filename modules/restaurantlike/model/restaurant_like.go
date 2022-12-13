package restaurantlikemodel

import (
	"errors"
	"fmt"
	"food_delivery/common"
	"time"
)

const EntityName = "UserLikeRestaurant"

type Like struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

func (l *Like) GetRestaurantId() int {
	return l.RestaurantId
}

func (l *Like) GetUserId() int {
	return l.UserId
}

func ErrCannotLikeRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this restaurant"),
		fmt.Sprintf("ERR_CANNOT_LIKE_RESTAURANT"),
	)
}

func ErrAlreadyLikedRestaurant() *common.AppError {
	return common.NewCustomError(
		errors.New("Already liked this restaurant"),
		fmt.Sprintf("Already liked this restaurant"),
		fmt.Sprintf("ERR_ALREADY_LIKED_RESTAURANT"),
	)
}

func ErrAlreadyUnLikedRestaurant() *common.AppError {
	return common.NewCustomError(
		errors.New("Already unliked this restaurant"),
		fmt.Sprintf("Already unliked this restaurant"),
		fmt.Sprintf("ERR_ALREADY_UNLIKED_RESTAURANT"),
	)
}
