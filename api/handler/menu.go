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

func (h *Handler) FetchMeals(c *gin.Context) {
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

	meals, err := h.MenuClient.FetchMeals(ctx, &pb.Filter{
		RestaurantId: restID,
		Limit:        int32(limit),
		Offset:       int32(offset),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error fetching meals from menu").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"Meals": meals})
}
