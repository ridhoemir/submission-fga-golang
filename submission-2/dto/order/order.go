package order

import (
	"submission-2/core"
	"time"
)

type OrderCreateRequest struct {
	OrderedAt    time.Time           `json:"orderedAt" binding:"required"`
	CustomerName string              `json:"customerName" binding:"required"`
	Items        []ItemCreateRequest `json:"items" binding:"required"`
}

func (req *OrderCreateRequest) ToOrder() core.Order {
	items := make([]core.Item, 0)
	for _, item := range req.Items {
		items = append(items, core.Item{
			Code:        item.Code,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}
	return core.Order{
		CustomerName: req.CustomerName,
		OrderedAt:    req.OrderedAt,
		Items:        items,
	}
}

type ItemCreateRequest struct {
	Code        string `json:"itemCode" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    int    `json:"quantity" binding:"required"`
}
