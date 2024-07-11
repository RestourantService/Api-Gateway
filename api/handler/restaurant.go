package handler

import (
	pb "api-gateway/genproto/restaurant"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (h *Handler) CreateRestaurant(c *gin.Context) {
	var rest pb.RestaurantDetails
	err := c.ShouldBind(&rest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	id, err := h.RestaurantClient.CreateRestaurant(ctx, &rest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error creating restaurant").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"New restaurant id": id.Id})
}

func (h *Handler) GetRestaurantByID(c *gin.Context) {
	id := c.Param("restaurant_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid restaurant id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	rest, err := h.RestaurantClient.GetRestaurantByID(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error getting restaurant").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Restaurant": rest})
}

func (h *Handler) UpdateRestaurant(c *gin.Context) {
	id := c.Param("restaurant_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid restaurant id").Error()})
		log.Println(err)
		return
	}

	var rest pb.RestaurantInfo
	err = c.ShouldBind(&rest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}
	rest.Id = id

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.RestaurantClient.UpdateRestaurant(ctx, &rest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error updating restaurant").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "Restaurant updated successfully")
}

func (h *Handler) DeleteRestaurant(c *gin.Context) {
	id := c.Param("restaurant_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid restaurant id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.RestaurantClient.DeleteRestaurant(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error deleting restaurant").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "Restaurant deleted successfully")
}

func (h *Handler) FetchRestaurants(c *gin.Context) {
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

	rests, err := h.RestaurantClient.FetchRestaurants(ctx, &pb.Pagination{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error fetching restaurants").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Restaurants": rests})
}
