package pkg

import (
	"api-gateway/config"
	pbm "api-gateway/genproto/menu"
	pbp "api-gateway/genproto/payment"
	pbreser "api-gateway/genproto/reservation"
	pbr "api-gateway/genproto/restaurant"
	pbu "api-gateway/genproto/user"
	"log"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(cfg *config.Config) pbu.UserClient {
	conn, err := grpc.NewClient(cfg.USER_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbu.NewUserClient(conn)
}

func NewRestaurantClient(cfg *config.Config) pbr.RestaurantClient {
	conn, err := grpc.NewClient(cfg.RESERVATION_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbr.NewRestaurantClient(conn)
}

func NewReservationClient(cfg *config.Config) pbreser.ReservationClient {
	conn, err := grpc.NewClient(cfg.RESERVATION_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbreser.NewReservationClient(conn)
}

func NewMenuClient(cfg *config.Config) pbm.MenuClient {
	conn, err := grpc.NewClient(cfg.RESERVATION_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbm.NewMenuClient(conn)
}

func NewPaymentClient(cfg *config.Config) pbp.PaymentClient {
	conn, err := grpc.NewClient(cfg.PAYMENT_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(errors.Wrap(err, "failed to connect to the address"))
		return nil
	}

	return pbp.NewPaymentClient(conn)
}
