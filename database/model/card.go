package model

type ProductCard struct {
	ID        uint `gorm:"primary_key;auto_increment"`
	ProductID uint `gorm:"unique_index:idx_user_id_product_id"`
	UserID    uint `gorm:"unique_index:idx_user_id_product_id"`
	Quantity  uint
	Product   *Product `gorm:"foreignKey:ProductID"`
	User      User     `gorm:"foreignKey:UserID"`
}

func (ProductCard) TableName() string {
	return "product_cards"
}
