package restaurantbiz

import (
	"context"
	"food_delivery/modules/restaurant/restaurantmodel"
)

//type Animal interface {
//	Run(ctx context.Context) error
//	Sleep(ctx context.Context) error
//}
//
//type Dog struct {
//	name string
//}
//
//func (d Dog) Run(ctx context.Context) error {
//	fmt.Println(d.name)
//	return nil
//}
//
//func (d Dog) Sleep(ctx context.Context) error {
//	fmt.Println(d.name)
//	return nil
//}
// router => handler => controller => services => repository
//router => handler (json)
//1. Validate
//2. Json data/ Body => create model ABC
//json from handler
//ABC {
//	name = json.Name
//	age = json.Age
//}
//
//if from == "" => return error
//
//=> controller/services
//Nhan ABC tu handler , xu li logic,
//
//if from > today => error
//tu ABC => parse model ABC'
//
//=> repository

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

// private
type createRestaurantBiz struct {
	store CreateRestaurantStore
}

// public | export for outside to use
func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

// one of method `createRestaurantBiz`
func (biz *createRestaurantBiz) CreateRestaurant(
	ctx context.Context,
	data *restaurantmodel.RestaurantCreate) error {

	if err := data.Validate(); err != nil {
		return err
	}
	err := biz.store.Create(ctx, data)

	return err
}
