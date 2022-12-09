package subscriber

import (
	"context"
	"food_delivery/component/appctx"
)

func Setup(appCtx appctx.AppContext) {
	IncreaseLikeCountAfterUserLikeRestaurant(appCtx, context.Background())
}
