package handler

import (
	"github.com/alvingxv/Assignment2-Hacktiv8-Golang/database"
	"github.com/alvingxv/Assignment2-Hacktiv8-Golang/repository/item_repository/item_pg"
	"github.com/alvingxv/Assignment2-Hacktiv8-Golang/repository/order_repository/order_pg"
	"github.com/alvingxv/Assignment2-Hacktiv8-Golang/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	r := gin.Default()

	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	itemRepo := item_pg.NewItemPG(db)
	orderRepo := order_pg.NewOrderPG(db)

	itemService := service.NewItemService(itemRepo)
	orderService := service.NewOrderService(orderRepo, itemService)

	orderHandler := NewOrderHandler(orderService)

	orderRoute := r.Group("/orders")
	{
		orderRoute.GET("/", orderHandler.GetAllOrder)
		orderRoute.POST("/", orderHandler.CreateOrder)
		orderRoute.DELETE("/:id", orderHandler.DeleteOrder)
		orderRoute.PUT("/:id", orderHandler.UpdateOrder)
	}

	r.Run("127.0.0.1:3030")
}
