package handler

import (
	pb "api-gateway/genproto/payment"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// CreatePayment godoc
// @Summary Creates a payment
// @Description Inserts new payment info to payments table in PostgreSQL
// @Tags payment
// @Param new_data body payment.PaymentDetails true "New data"
// @Success 200 {object} payment.Status
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while creating payment"
// @Router /reservation-system/payments [post]
func (h *Handler) CreatePayment(c *gin.Context) {
	var pay pb.PaymentDetails
	err := c.ShouldBind(&pay)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	status, err := h.PaymentClient.CreatePayment(ctx, &pay)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error creating payment").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Status": status.Status})
}

// GetPayment godoc
// @Summary Gets a payment
// @Description Retrieves payment info from payments table in PostgreSQL
// @Tags payment
// @Param payment_id path string true "Payment ID"
// @Success 200 {object} payment.PaymentInfo
// @Failure 400 {object} string "Invalid payment ID"
// @Failure 500 {object} string "Server error while getting payment"
// @Router /reservation-system/payments/{payment_id} [get]
func (h *Handler) GetPayment(c *gin.Context) {
	id := c.Param("payment_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid payment id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	pay, err := h.PaymentClient.GetPayment(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error getting payment").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Payment": pay})
}

// UpdatePayment godoc
// @Summary Updates a payment
// @Description Updates payment info in payments table in PostgreSQL
// @Tags payment
// @Accept json
// @Produce json
// @Param payment_id path string true "Payment ID"
// @Param new_info body payment.PaymentInsert true "New info"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid payment ID or data"
// @Failure 500 {object} string "Server error while updating payment"
// @Router /reservation-system/payments/{payment_id} [put]
func (h *Handler) UpdatePayment(c *gin.Context) {
	id := c.Param("payment_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid payment id").Error()})
		log.Println(err)
		return
	}

	var pay pb.PaymentInfo
	err = c.ShouldBind(&pay)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}
	pay.Id = id

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.PaymentClient.UpdatePayment(ctx, &pay)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error updating payment").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, "Payment updated successfully")
}
