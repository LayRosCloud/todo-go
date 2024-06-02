package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/layroscloud/todo-go/pkg/dto"
	"net/http"
	"strconv"
)

func (h *Handler) FindLists(c *gin.Context) {
	id, _ := c.Get("userId")

	list, errQuery := h.services.FindAllLists(id.(int64))
	if errQuery != nil {
		NewErrorResponse(c, http.StatusInternalServerError, errQuery.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) FindByIdList(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	list, errQuery := h.services.FindByIdList(id)
	if errQuery != nil {
		NewErrorResponse(c, http.StatusNotFound, errQuery.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          list.Id,
		"title":       list.Title,
		"description": list.Description,
	})
}

func (h *Handler) CreateList(c *gin.Context) {
	idStr, _ := c.Get("userId")
	var input dto.ListCreateDto

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "list has bad fields: "+err.Error())
		return
	}
	id, err := h.services.CreateList(input, idStr.(int64))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) UpdateList(c *gin.Context) {
	var input dto.ListUpdateDto

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "list has bad fields")
		return
	}
	id, err := h.services.UpdateList(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) DeleteList(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.ParseInt(idString, 10, 64)
	status, err := h.services.DeleteList(id)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})

}
