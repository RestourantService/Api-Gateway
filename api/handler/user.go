package handler

import (
	pb "api-gateway/genproto/user"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// GetUser godoc
// @Summary Gets a user
// @Description Retrieves a user from users table in PostgreSQL
// @Tags user
// @Accept json
// @Produce json
// @Param user_id path uuid true "User ID"
// @Success 200 {object} user.UserInfo
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Server error getting user"
// @Router /:user_id [get]
func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid user id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	user, err := h.UserClient.GetUser(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error getting user").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"User": user})
}

// UpdateUser godoc
// @Summary Updates a user
// @Description Updates a user in users table in PostgreSQL
// @Tags user
// @Accept json
// @Produce json
// @Param user_id path uuid true "User ID"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Server error updating user"
// @Router /:user_id [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid user id").Error()})
		log.Println(err)
		return
	}

	var user pb.UserInfo
	err = c.ShouldBind(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid data").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.UserClient.UpdateUser(ctx, &user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error updating user").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, "User updated successfully")
}

// DeleteUser godoc
// @Summary Deletes a user
// @Description Deletes a user from users table in PostgreSQL
// @Tags user
// @Accept json
// @Produce json
// @Param user_id path uuid true "User ID"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Server error deleting user"
// @Router /:user_id [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("user_id")
	_, err := uuid.Parse(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": errors.Wrap(err, "invalid user id").Error()})
		log.Println(err)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	_, err = h.UserClient.DeleteUser(ctx, &pb.ID{Id: id})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": errors.Wrap(err, "error deleting user").Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, "User deleted successfully")
}
