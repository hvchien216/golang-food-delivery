package main

import (
	"fmt"
	"food_delivery/component/appctx"
	"food_delivery/modules/restaurant/restaurantmodel"
	"food_delivery/modules/restaurant/restauranttransport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	//refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(db, err)

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	appCtx := appctx.NewAppContext(db)

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

		restaurants.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			var data restaurantmodel.Restaurant

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(401, gin.H{"error": err.Error()})

				return
			}

			c.JSON(http.StatusOK, data)
		})

		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

		restaurants.PATCH("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			var data restaurantmodel.RestaurantUpdate

			if err := c.ShouldBind(&data); err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
				c.JSON(401, gin.H{"error": err.Error()})

				return
			}

			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})

		restaurants.DELETE("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))

			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			if err := db.Table(restaurantmodel.Restaurant{}.TableName()).
				Where("id = ?", id).
				Delete(nil).Error; err != nil {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{"ok": 1})
		})
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
