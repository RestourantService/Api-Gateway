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

// CreateReservation godoc
// @Summary Creates a reservation
// @Description Inserts new reservation info to reservations table in PostgreSQL
// @Tags reservation
// @Param new_data body reservation.ReservationDetails true "New data"
// @Success 200 {object} json
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while creating reservation"
// @Router /reservation-system/reservations [post]
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

// GetReservationByID godoc
// @Summary Gets a reservation
// @Description Retrieves reservation info from reservations table in PostgreSQL
// @Tags reservation
// @Param reservation_id path string true "Reservation ID"
// @Success 200 {object} reservation.ReservationInfo
// @Failure 400 {object} string "Invalid reservation ID"
// @Failure 500 {object} string "Server error while getting reservation"
// @Router /reservation-system/reservations/{reservation_id} [get]
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

// UpdateReservation godoc
// @Summary Updates a reservation
// @Description Updates reservation info in reservations table in PostgreSQL
// @Tags reservation
// @Accept json
// @Produce json
// @Param reservation_id path string true "Reservation ID"
// @Param new_info body reservation.ReservationInfo true "New info"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid reservation ID or data"
// @Failure 500 {object} string "Server error while updating reservation"
// @Router /reservation-system/reservations/{reservation_id} [put]
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

	c.JSON(http.StatusOK, "Reservation updated successfully")
}

// DeleteReservation godoc
// @Summary Deletes a reservation
// @Description Deletes reservation info from reservations table in PostgreSQL
// @Tags reservation
// @Param reservation_id path string true "Reservation ID"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid reservation ID"
// @Failure 500 {object} string "Server error while deleting reservation"
// @Router /reservation-system/reservations/{reservation_id} [delete]
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

	c.JSON(http.StatusOK, "Reservation deleted successfully")
}

// ValidateReservation godoc
// @Summary Validates a reservation
// @Description Checks whether a reservation exists in reservations table in PostgreSQL
// @Tags reservation
// @Param reservation_id path string true "Reservation ID"
// @Success 200 {object} reservation.Status
// @Failure 400 {object} string "Invalid reservation ID"
// @Failure 500 {object} string "Server error while checking reservation"
// @Router /reservation-system/{reservation_id}/check [get]
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

// Order godoc
// @Summary Orders meals
// @Description Inserts order for a reservation in Redis
// @Tags reservation
// @Param reservation_id path string true "Reservation ID"
// @Param order body reservation.ReservationOrders true "New order"
// @Success 200 {object} reservation.ID
// @Failure 400 {object} string "Invalid reservation ID or data"
// @Failure 500 {object} string "Server error while ordering"
// @Router /reservation-system/{reservation_id}/order [post]
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

// Pay godoc
// @Summary Pays for a reservation
// @Description Inserts payment info to payments table in PostgreSQL
// @Tags reservation
// @Param reservation_id path string true "Reservation ID"
// @Success 200 {object} reservation.Status
// @Failure 400 {object} string "Invalid reservation ID"
// @Failure 500 {object} string "Server error while making a payment"
// @Router /reservation-system/{reservation_id}/payment [post]
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

// FetchReservations godoc
// @Summary Fetches reservations
// @Description Retrieves multiple reservations info from reservations table in PostgreSQL
// @Tags reservation
// @Param user_id query string false "User ID"
// @Param restaurant_id query string false "Restaurant ID"
// @Param limit path string false "Number of reservations to fetch"
// @Param offset path string false "Number of reservations to omit"
// @Success 200 {object} reservation.Reservations
// @Failure 400 {object} string "Invalid pagination parameters"
// @Failure 500 {object} string "Server error while fetching reservations"
// @Router /reservation-system/reservations [get]
func (h *Handler) FetchReservations(c *gin.Context) {
	filter := pb.Filter{
		UserId:       c.Query("user_id"),
		RestaurantId: c.Query("restaurant_id"),
	}
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": errors.Wrap(err, "invalid pagination parameters").Error()})
			log.Println(err)
			return
		}
		filter.Limit = int32(limit)
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": errors.Wrap(err, "invalid pagination parameters").Error()})
			log.Println(err)
			return
		}
		filter.Offset = int32(offset)
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	resers, err := h.ReservationClient.FetchReservations(ctx, &filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error fetching reservations").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Reservations": resers})
}
