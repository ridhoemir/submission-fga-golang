package core

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"-"`
	Code        string `gorm:"type:varchar(10)" json:"itemCode"`
	Description string `gorm:"type:varchar(50)" json:"description"`
	Quantity    int    `gorm:"type:int" json:"quantity"`
	OrderID     uint   `gorm:"foreignKey:OrderID" json:"-"`
}
