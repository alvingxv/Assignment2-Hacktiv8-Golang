package order_repository

import (
	"github.com/alvingxv/Assignment2-Hacktiv8-Golang/entity"
	"github.com/alvingxv/Assignment2-Hacktiv8-Golang/pkg/errs"
)

type OrderRepository interface {
	CreateOrder(orderPayload entity.Order, itemsPayload []entity.Item) (*entity.Order, errs.MessageErr)
	GetAllOrder() (*[]OrderItem, errs.MessageErr)
	UpdateOrder(orderPayload entity.Order, itemsPayload []entity.Item) (*OrderItem, errs.MessageErr)
	DeleteOrder(orderId int) errs.MessageErr
}
