package common

import (
	driver "database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url"`
	Width     int    `json:"width" gorm:"column:width"`
	Height    int    `json:"height" gorm:"column:height"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string {
	return "images"
}

func (img *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var image Image
	if err := json.Unmarshal(bytes, &image); err != nil {
		return err
	}

	// syntax: *pointer_variable = value | when we change value of a pointer
	*img = image
	return nil
}

// write down DB, must be implement method Value
func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}

type Images []Image

func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if ok {
		return errors.New(fmt.Sprintf("Failed to unmarshal JSONB %s", value))
	}

	var image []Image
	if err := json.Unmarshal(bytes, &image); err != nil {
		return err
	}

	*j = image
	return nil
}

func (j *Images) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}

	return json.Marshal(j)
}
