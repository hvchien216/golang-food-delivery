package restaurantmodel

import "testing"

type DataTable struct {
	Input  RestaurantCreate
	Expect error
}

func TestRestaurantCreate_Validate(t *testing.T) {
	table := []DataTable{
		{Input: RestaurantCreate{Name: ""}, Expect: ErrNameCannotEmpty},
		{Input: RestaurantCreate{Name: "Hello there"}, Expect: nil},
	}

	for _, d := range table {
		err := d.Input.Validate()

		if err != d.Expect {
			t.Errorf("Test Validate() failed, expected %e with %#v, but got  %v", d.Expect, d.Input, err)
			break
		}
	}

	t.Log("Test Validate() passed")
}
