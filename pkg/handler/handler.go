package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/layroscloud/todo-go/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		services: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/v1/accounts")
	{
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/sign-up", h.SignUp)
	}

	api := router.Group("/api", h.UserIdentity)
	{
		lists := api.Group("v1/lists")
		{
			lists.GET("/", h.FindLists)
			lists.GET("/:id", h.FindByIdList)
			lists.POST("/", h.CreateList)
			lists.PUT("/:id", h.UpdateList)
			lists.DELETE("/:id", h.DeleteList)
		}

		items := lists.Group("/:id/items")
		{
			items.GET("/", h.FindItems)
			items.POST("/", h.CreateItem)
			items.GET("/:item_id", h.FindByIdItem)
			items.PUT("/:item_id", h.UpdateItem)
			items.DELETE("/:item_id", h.DeleteItem)
		}
	}

	return router
}
