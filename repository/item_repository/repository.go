package item_repository

import (
	"github.com/alvingxv/Assignment2-Hacktiv8-Golang/entity"
	"github.com/alvingxv/Assignment2-Hacktiv8-Golang/pkg/errs"
)

type ItemRepository interface {
	FindItemsByItemCodes(itemCodes []string) ([]*entity.Item, errs.MessageErr)
}
