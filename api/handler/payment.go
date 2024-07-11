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
