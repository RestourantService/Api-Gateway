package handler

import (
	"api-gateway/config"
	"api-gateway/genproto/menu"
	"api-gateway/genproto/payment"
	"api-gateway/genproto/reservation"
	"api-gateway/genproto/restaurant"
	"api-gateway/genproto/user"
	"api-gateway/pkg"
)

type Handler struct {
	UserClient        user.UserClient
	RestaurantClient  restaurant.RestaurantClient
	ReservationClient reservation.ReservationClient
	MenuClient        menu.MenuClient
	PaymentClient     payment.PaymentClient
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		UserClient:        pkg.NewUserClient(cfg),
		RestaurantClient:  pkg.NewRestaurantClient(cfg),
		ReservationClient: pkg.NewReservationClient(cfg),
		MenuClient:        pkg.NewMenuClient(cfg),
		PaymentClient:     pkg.NewPaymentClient(cfg),
	}
}
