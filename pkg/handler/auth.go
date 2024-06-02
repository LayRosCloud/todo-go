package handler

import (
	"github.com/layroscloud/todo-go/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/layroscloud/todo-go/entity"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, errQuery := h.services.Authorization.CreateUser(input)
	if errQuery != nil {
		NewErrorResponse(c, http.StatusInternalServerError, errQuery.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input dto.SignInDto

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
