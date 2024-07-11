package handler

import (
	pb "api-gateway/genproto/reservation"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (h *Handler) CreateReservation(c *gin.Context) {
	var res pb.ReservationDetails
	err := c.ShouldBind(&res)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	id, err := h.ReservationClient.CreateReservation(ctx, &res)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error creating reservation").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"New reservation id": id.Id})
}

func (h *Handler) GetReservationByID(c *gin.Context) {
	id := c.Param("reservation_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid reservation id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	res, err := h.ReservationClient.GetReservationByID(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error getting reservation").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Reservation": res})
}

func (h *Handler) UpdateReservation(c *gin.Context) {
	id := c.Param("reservation_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid reservation id").Error()})
		log.Println(err)
		return
	}

	var res pb.ReservationInfo
	err = c.ShouldBind(&res)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}
	res.Id = id

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.ReservationClient.UpdateReservation(ctx, &res)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error updating reservation").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "Reservation updated successfully")
}

func (h *Handler) DeleteReservation(c *gin.Context) {
	id := c.Param("reservation_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid reservation id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.ReservationClient.DeleteReservation(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error deleting reservation").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "Reservation deleted successfully")
}

func (h *Handler) ValidateReservation(c *gin.Context) {
	id := c.Param("reservation_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid reservation id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	status, err := h.ReservationClient.ValidateReservation(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error checking reservation").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Valid reservation": status.Successful})
}

func (h *Handler) Order(c *gin.Context) {
	id := c.Param("reservation_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid reservation id").Error()})
		log.Println(err)
		return
	}

	var resOrd pb.ReservationOrders
	err = c.ShouldBind(&resOrd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}
	resOrd.Id = id

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	ordID, err := h.ReservationClient.Order(ctx, &resOrd)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error ordering").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"New order id": ordID.Id})
}

func (h *Handler) Pay(c *gin.Context) {
	id := c.Param("reservation_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid reservation id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	status, err := h.ReservationClient.Pay(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error making a payment").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Payment success": status.Successful})
}

func (h *Handler) FetchReservations(c *gin.Context) {
	userID := c.Query("user_id")
	restID := c.Query("restaurant_id")

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid pagination parameters").Error()})
		log.Println(err)
		return
	}
	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid pagination parameters").Error()})
		log.Println(err)
		return
	}
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	resers, err := h.ReservationClient.FetchReservations(ctx, &pb.Filter{
		UserId:       userID,
		RestaurantId: restID,
		Limit:        int32(limit),
		Offset:       int32(offset),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error fetching reservations").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Reservations": resers})
}
