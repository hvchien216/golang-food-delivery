package restaurantlikemodel

// Filter : TODO
// for both case:
// Get Users like restaurant
// Get restaurants are liked by user
type Filter struct {
	RestaurantId int `json:"-" form:"restaurant_id"`
	UserId       int `json:"-" form:"user_id"`
}
