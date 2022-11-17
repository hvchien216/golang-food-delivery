package ginupload

import (
	"fmt"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"github.com/gin-gonic/gin"
)

// Upload file to S3
// 1. Get image/file from request header
// 2. Check file is real image
// 3. Save image
// 1. Save to local machine
// 2. Save to cloud storage (S3)
// 3. Improve security

func Upload(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))
		c.JSON(200, common.SimpleSucessResponse(true))
	}
}
