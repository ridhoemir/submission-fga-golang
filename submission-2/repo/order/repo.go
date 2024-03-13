package order

import (
	"errors"
	"submission-2/core"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateOrder(order core.Order) (core.Order, error) {
	err := r.db.Debug().Create(&order).Error

	return order, err

}

func (r *Repository) GetAllOrders() ([]core.Order, error) {
	var orders []core.Order

	err := r.db.Debug().Preload("Items").Find(&orders).Error

	return orders, err
}

func (r *Repository) UpdateOrder(order core.Order, id uint64) (core.Order, error) {
	var fetchedOrder core.Order
	var items []core.Item
	err := r.db.Preload("Items").First(&fetchedOrder, id).Error
	if err != nil {
		return fetchedOrder, err
	}

	result := r.db.Where("order_id = ?", id).Find(&items)
	if result.RowsAffected == 0 {
		return fetchedOrder, errors.New("item not found")
	}

	for i, item := range items {
		err = r.db.Model(&item).Updates(core.Item{
			Code:        order.Items[i].Code,
			Description: order.Items[i].Description,
			Quantity:    order.Items[i].Quantity,
		}).Error
		if err != nil {
			return fetchedOrder, err
		}
	}

	err =
		r.db.Where("id = ?", id).Model(&fetchedOrder).Updates(core.Order{
			CustomerName: order.CustomerName,
			OrderedAt:    order.OrderedAt,
		}).Error

	order.ID = uint(id)
	return order, err
}

func (r *Repository) DeleteOrder(id uint64) (string, error) {
	var order core.Order
	var items []core.Item
	var message string

	err := r.db.Where("id= ?", id).First(&order).Error
	if err != nil {
		message = "Order not found "
		return message, err
	}
	result := r.db.Where("order_id = ?", id).Find(&items)
	if result.RowsAffected == 0 {
		message = "Item not found "
		return message, result.Error
	}

	for _, item := range items {
		err = r.db.Delete(&item).Error
		if err != nil {
			message = "Error Delete Item"
			return message, err
		}
	}
	err = r.db.Delete(&order).Error
	if err != nil {
		message = "Error Delete Order"
		return message, err
	}

	message = "Success Delete"

	return message, err

}
