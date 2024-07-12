package handler

import (
	pb "api-gateway/genproto/menu"
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// AddMeal godoc
// @Summary Adds a meal to menu
// @Description Inserts new meal info to menu table in PostgreSQL
// @Tags menu
// @Param new_data body menu.MealDetails true "New data"
// @Success 200 {object} menu.ID
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while adding meal to menu"
// @Router /reservation-system/menu [post]
func (h *Handler) AddMeal(c *gin.Context) {
	var meal pb.MealDetails
	err := c.ShouldBind(&meal)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	id, err := h.MenuClient.AddMeal(ctx, &meal)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error adding meal to menu").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"New meal id": id.Id})
}

// GetMealByID godoc
// @Summary Gets a meal
// @Description Retrieves meal info from menu table in PostgreSQL
// @Tags menu
// @Param meal_id path string true "Meal ID"
// @Success 200 {object} menu.MealInfo
// @Failure 400 {object} string "Invalid meal ID"
// @Failure 500 {object} string "Server error while getting meal from menu"
// @Router /reservation-system/menu/{meal_id} [get]
func (h *Handler) GetMealByID(c *gin.Context) {
	id := c.Param("meal_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid meal id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	meal, err := h.MenuClient.GetMealByID(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error getting meal from menu").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Meal": meal})
}

// UpdateMeal godoc
// @Summary Updates a meal
// @Description Updates meal info in menu table in PostgreSQL
// @Tags menu
// @Accept json
// @Produce json
// @Param meal_id path string true "Meal ID"
// @Param new_info body meal.MealInfo true "New info"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid meal ID or data"
// @Failure 500 {object} string "Server error while updating meal in menu"
// @Router /reservation-system/menu/{meal_id} [put]
func (h *Handler) UpdateMeal(c *gin.Context) {
	id := c.Param("meal_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid meal id").Error()})
		log.Println(err)
		return
	}

	var meal pb.MealInfo
	err = c.ShouldBind(&meal)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}
	meal.Id = id

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.MenuClient.UpdateMeal(ctx, &meal)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error updating meal in menu").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "Meal updated successfully")
}

// DeleteMeal godoc
// @Summary Deletes a meal
// @Description Removes meal info from menu table in PostgreSQL
// @Tags menu
// @Param meal_id path string true "Meal ID"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid meal ID"
// @Failure 500 {object} string "Server error while removing meal from menu"
// @Router /reservation-system/menu/{meal_id} [delete]
func (h *Handler) DeleteMeal(c *gin.Context) {
	id := c.Param("meal_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid meal id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.MenuClient.DeleteMeal(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error removing meal from menu").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "Meal removed successfully")
}

// FetchMeals godoc
// @Summary Fetches meals
// @Description Retrieves multiple meals info from menu table in PostgreSQL
// @Tags menu
// @Param restaurant_id query string false "Restaurant ID"
// @Param limit path string false "Number of meals to fetch"
// @Param offset path string false "Number of meals to omit"
// @Success 200 {object} menu.Meals
// @Failure 400 {object} string "Invalid pagination parameters"
// @Failure 500 {object} string "Server error while fetching meals from menu"
// @Router /reservation-system/menu [get]
func (h *Handler) FetchMeals(c *gin.Context) {
	filter := pb.Filter{
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

	meals, err := h.MenuClient.FetchMeals(ctx, &filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error fetching meals from menu").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Meals": meals})
}
