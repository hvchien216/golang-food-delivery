package main

import (
	"fmt"
	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/uploadprovider"
	"food_delivery/middleware"
	"food_delivery/modules/upload/uploadtransport/ginupload"
	"food_delivery/modules/user/usertransport/ginuser"
	"food_delivery/pubsub/pblocal"
	"food_delivery/routes/restaurantroute"
	"food_delivery/skio"
	"food_delivery/subscriber"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	//refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DBConnectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(db, err)
	db = db.Debug()

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, provider uploadprovider.UploadProvider, secretKey string) error {
	appCtx := appctx.NewAppContext(db, provider, secretKey, pblocal.NewPubsub())

	r := gin.Default()

	rtEngine := skio.NewEngine()

	if err := rtEngine.Run(appCtx, r); err != nil {
		log.Fatalln(err)
	}

	//deprecated
	//subscriber.Setup(appCtx)

	// use this line as an alternative for Setup
	if err := subscriber.NewEngine(appCtx, rtEngine).Start(); err != nil {
		log.Fatalln()
	}

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.StaticFile("/demo/", "./demo.html")

	v1 := r.Group("/v1")
	v1.POST("/upload", ginupload.Upload(appCtx))
	//v1.GET("/presigned-upload-url", func(C *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"data": s3Provider.GetUploadPresignedUrl(c.Request.Context())})
	//})

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfile(appCtx))

	v1.GET("/encode-uid", func(c *gin.Context) {
		type reqData struct {
			DbType int `form:"type"`
			RealId int `form:"id"`
		}

		var d reqData
		c.ShouldBind(&d)

		c.JSON(http.StatusOK, gin.H{"id": common.NewUID(uint32(d.RealId), d.DbType, 1)})
	})
	restaurantroute.Routes(v1, appCtx)

	admin := v1.Group(
		"/admin",
		middleware.RequireAuth(appCtx),
		middleware.RequireRoles(appCtx, "admin"),
	)
	{
		admin.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, common.SimpleSucessResponse("ok"))
		})
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//type Permission uint64
//
//const (
//	Read Permission = 1 << iota
//	Write
//	Delete
//	Invite
//)

// User 1: Create group FB A (Owner A)
// User 1 invite User 2 into group A (User 2 as a member of A)

// user_id | group_id | permission
// 1	   | 1        | 6

//or

// user_id | group_id | permission
// 1	   | 1        | read,write,delete,invite
