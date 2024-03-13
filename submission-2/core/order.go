package core

import "time"

type Order struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerName string    `gorm:"type:varchar(50)" json:"customerName"`
	OrderedAt    time.Time `gorm:"type:timestamp" json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:OrderID" json:"items"`
}
