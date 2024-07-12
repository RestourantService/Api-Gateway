package handler

import (
	_ "api-gateway/genproto/authentication"
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
// @Description Retrieves user info from users table in PostgreSQL
// @Tags user
// @Param user_id path string true "User ID"
// @Success 200 {object} user.UserInfo
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Server error while getting user"
// @Router /reservation-system/users/{user_id} [get]
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
// @Description Updates user info in users table in PostgreSQL
// @Tags user
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param new_info body authentication.UserDetails true "New info"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid user ID or data"
// @Failure 500 {object} string "Server error while updating user"
// @Router /reservation-system/users/{user_id} [put]
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
	user.Id = id

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
// @Description Deletes user info from users table in PostgreSQL
// @Tags user
// @Param user_id path string true "User ID"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid user ID"
// @Failure 500 {object} string "Server error while deleting user"
// @Router /reservation-system/users/{user_id} [delete]
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
