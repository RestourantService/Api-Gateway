package api

import (
	"api-gateway/api/handler"
	// "api-gateway/api/middleware"
	"api-gateway/config"

	_ "api-gateway/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Restaurant Reservation System
// @version 1.0
// @description API Gateway of Restaurant Reservation System
// @host localhost:8080
// BasePath: /
func NewRouter(cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/reservation-system")
	// api.Use(middleware.Check)

	h := handler.NewHandler(cfg)

	u := api.Group("/users")
	{
		u.GET("/:user_id", h.GetUser)
		u.PUT("/:user_id", h.UpdateUser)
		u.DELETE("/:user_id", h.DeleteUser)
	}

	rest := api.Group("/restaurants")
	{
		rest.POST("", h.CreateRestaurant)
		rest.GET("/:restaurant_id", h.GetRestaurantByID)
		rest.PUT("/:restaurant_id", h.UpdateRestaurant)
		rest.DELETE("/:restaurant_id", h.DeleteRestaurant)
		rest.GET("", h.FetchRestaurants)
	}

	reser := api.Group("reservations")
	{
		reser.POST("", h.CreateReservation)
		reser.GET("/:reservation_id", h.GetReservationByID)
		reser.PUT("/:reservation_id", h.UpdateReservation)
		reser.DELETE("/:reservation_id", h.DeleteReservation)
		reser.GET("/:reservation_id/check", h.ValidateReservation)
		reser.POST("/:reservation_id/order", h.Order)
		reser.POST("/:reservation_id/payment", h.Pay)
		reser.GET("", h.FetchReservations)
	}

	m := api.Group("menu")
	{
		m.POST("", h.AddMeal)
		m.GET("/:meal_id", h.GetMealByID)
		m.PUT("/:meal_id", h.UpdateMeal)
		m.DELETE("/:meal_id", h.DeleteMeal)
		m.GET("", h.FetchMeals)
	}

	p := api.Group("payments")
	{
		p.POST("", h.CreatePayment)
		p.GET("/:payment_id", h.GetPayment)
		p.PUT("/:payment_id", h.UpdatePayment)
	}

	return r
}
