package ginuser

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/hasher"
	"food_delivery/modules/user/userbiz"
	"food_delivery/modules/user/usermodel"
	"food_delivery/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		repo := userbiz.NewRegisterBusiness(store, md5)

		if err := repo.Register(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.FakeId.String()))
	}
}
