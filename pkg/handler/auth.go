package handler

import (
	"net/http"

	"github.com/ShawaDev/auth/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	id, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	token, role, err := h.service.GenerateToken(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"role":  role,
	})

}
