package restaurantmodel

import (
	"errors"
	"food_delivery/common"
	"strings"
)

const EntityName = "RESTAURANT"

var (
	ErrNameCannotEmpty = errors.New("restaurant name cannot be blank")
)

// Business Model
// `common.SQLModel` is embed struct
type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string             `json:"name" gorm:"column:name;"`
	Addr            string             `json:"address" gorm:"column:addr;"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images     `json:"cover" gorm:"column:cover;"`
	LikedCount      int                `json:"liked_count" gorm:"column:liked_count;"` // computed field
	UserId          int                `json:"-" gorm:"column:owner_id"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) GetOwnerId() int {
	return r.UserId
}

// Data Model
type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name;"`
	Addr  *string        `json:"address" gorm:"column:addr;"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

// Data Model
type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	OwnerId         int            `json:"-" gorm:"column:owner_id"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}
func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return ErrNameCannotEmpty
	}
	return nil
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)

	if u := r.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}
