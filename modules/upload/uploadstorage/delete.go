package uploadstorage

import (
	"context"
	"food_delivery/common"
)

func (store sqlStore) DeleteImages(ctx context.Context, ids []int) error {
	db := store.db

	if err := db.Table(common.Image{}.TableName()).
		Where("id in (?)", ids).
		Delete(nil).
		Error; err != nil {
		return err
	}

	return nil
}