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

// CreateRestaurant godoc
// @Summary Creates a restaurant
// @Description Inserts new restaurant info to restaurants table in PostgreSQL
// @Tags restaurant
// @Param new_data body restaurant.RestaurantDetails true "New data"
// @Success 200 {object} restaurant.ID
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while creating restaurant"
// @Router /reservation-system/restaurants [post]
func (h *Handler) CreateRestaurant(c *gin.Context) {
	h.Logger.Info("CreateRestaurant method is starting")

	var rest pb.RestaurantDetails
	err := c.ShouldBind(&rest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		
		err := errors.Wrap(err, "failed to data")
		h.Logger.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	id, err := h.RestaurantClient.CreateRestaurant(ctx, &rest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error creating restaurant").Error()})
		
		err := errors.Wrap(err, "failed to CreateRestaurant").Error()
		h.Logger.Error(err)
		return
	}

	h.Logger.Info("CreateRestaurant has successfully finished")
	c.JSON(http.StatusCreated, gin.H{"New restaurant id": id.Id})
}

// GetRestaurantByID godoc
// @Summary Gets a restaurant
// @Description Retrieves restaurant info from restaurants table in PostgreSQL
// @Tags restaurant
// @Param restaurant_id path string true "Restaurant ID"
// @Success 200 {object} restaurant.RestaurantInfo
// @Failure 400 {object} string "Invalid restaurant ID"
// @Failure 500 {object} string "Server error while getting restaurant"
// @Router /reservation-system/restaurants/{restaurant_id} [get]
func (h *Handler) GetRestaurantByID(c *gin.Context) {
	h.Logger.Info("GetRestaurantByID method is starting")

	id := c.Param("restaurant_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid restaurant id").Error()})
		
		err := errors.Wrap(err, "failed to data").Error()
		h.Logger.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	rest, err := h.RestaurantClient.GetRestaurantByID(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error getting restaurant").Error()})
		
		err := errors.Wrap(err, "failed to GetRestaurantByID").Error()
		h.Logger.Error(err)
		return
	}

	h.Logger.Info("GetRestaurantByID has successfully finished")
	c.JSON(http.StatusOK, gin.H{"Restaurant": rest})
}

// UpdateRestaurant godoc
// @Summary Updates a restaurant
// @Description Updates restaurant info in restaurants table in PostgreSQL
// @Tags restaurant
// @Accept json
// @Produce json
// @Param restaurant_id path string true "Restaurant ID"
// @Param new_info body restaurant.RestaurantInfo true "New info"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid restaurant ID or data"
// @Failure 500 {object} string "Server error while updating restaurant"
// @Router /reservation-system/restaurants/{restaurant_id} [put]
func (h *Handler) UpdateRestaurant(c *gin.Context) {
	h.Logger.Info("UpdateRestaurant method is starting")

	id := c.Param("restaurant_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid restaurant id").Error()})
		
		err := errors.Wrap(err, "failed to data").Error()
		h.Logger.Error(err)
		return
	}

	var rest pb.RestaurantInfo
	err = c.ShouldBind(&rest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		
		err := errors.Wrap(err, "failed to data").Error()
		h.Logger.Error(err)
		return
	}
	rest.Id = id

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.RestaurantClient.UpdateRestaurant(ctx, &rest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error updating restaurant").Error()})
		
		err := errors.Wrap(err, "failed to UpdateRestaurant").Error()
		h.Logger.Error(err)
		return
	}

	h.Logger.Info("UpdateRestaurant has successfully finished")
	c.JSON(http.StatusNoContent, "Restaurant updated successfully")
}

// DeleteRestaurant godoc
// @Summary Deletes a restaurant
// @Description Deletes restaurant info from restaurants table in PostgreSQL
// @Tags restaurant
// @Param restaurant_id path string true "Restaurant ID"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid restaurant ID"
// @Failure 500 {object} string "Server error while deleting restaurant"
// @Router /reservation-system/restaurants/{restaurant_id} [delete]
func (h *Handler) DeleteRestaurant(c *gin.Context) {
	h.Logger.Info("DeleteRestaurant method is starting")

	id := c.Param("restaurant_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid restaurant id").Error()})
		
		err := errors.Wrap(err, "failed to data").Error()
		h.Logger.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.RestaurantClient.DeleteRestaurant(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error deleting restaurant").Error()})
		
		err := errors.Wrap(err, "failed to DeleteRestaurant").Error()
		h.Logger.Error(err)
		return
	}

	h.Logger.Info("DeleteRestaurant has successfully finished")
	c.JSON(http.StatusNoContent, "Restaurant deleted successfully")
}

// FetchRestaurants godoc
// @Summary Fetches restaurants
// @Description Retrieves multiple restaurants info from restaurants table in PostgreSQL
// @Tags restaurant
// @Param limit path string false "Number of restaurants to fetch"
// @Param offset path string false "Number of restaurants to omit"
// @Success 200 {object} restaurant.Restaurants
// @Failure 400 {object} string "Invalid pagination parameters"
// @Failure 500 {object} string "Server error while fetching restaurants"
// @Router /reservation-system/restaurants [get]
func (h *Handler) FetchRestaurants(c *gin.Context) {
	h.Logger.Info("FetchRestaurants method is starting")
	
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

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	rests, err := h.RestaurantClient.FetchRestaurants(ctx, &pb.Pagination{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error fetching restaurants").Error()})
		
		err := errors.Wrap(err, "failed to FetchRestaurants").Error()
		h.Logger.Error(err)
		return
	}

	h.Logger.Info("FetchRestaurants has successfully finished")
	c.JSON(http.StatusOK, gin.H{"Restaurants": rests})
}
