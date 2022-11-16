package restaurantmodel

type Filter struct {
	CityId int `json:"city_id,omitempty" form:"city_id"`
}
