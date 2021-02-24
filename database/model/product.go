package model

type Product struct {
	ID          uint   `gorm:"primary_key;auto_increment"`
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:1024"`
	Quantity    uint
	Price       float64
}

func (Product) TableName() string {
	return "products"
}
