package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/neutron2025/ecommerce-yt/controllers"
	"github.com/neutron2025/ecommerce-yt/database"
	"github.com/neutron2025/ecommerce-yt/middleware"
	"github.com/neutron2025/ecommerce-yt/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRouters(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":" + port))

}
