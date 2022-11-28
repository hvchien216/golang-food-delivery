package uploadprovider

import (
	"context"
	"food_delivery/common"
)

type UploadProvider interface {
	SaveFileUploaded(context context.Context, data []byte, dst string) (*common.Image, error)
}
