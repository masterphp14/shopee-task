package model

import "tax-calculator/engine"

// Item model
type Item struct {
	ID        int      `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	Order     *Order   `json:"order,omitempty"`
	OrderID   int      `gorm:"order_id" json:"order_id,omitempty"`
	ItemName  string   `gorm:"column:item_name" json:"item_name"`
	TaxCode   *Taxcode `json:"tax_code,omitempty"`
	TaxCodeID int      `gorm:"tax_code_id" json:"tax_code_id,omitempty"`
	Price     float64  `gorm:"column:price" json:"price"`
	Tax       float64  `gorm:"-" json:"tax"`
	Amount    float64  `gorm:"-" json:"amount"`
}

// TableName table name in database
func (m *Item) TableName() string {
	return "item"
}

// Read to get one data in item based on parameter (using preload to get foreign key; not efficient)
func (m *Item) Read(condition string, param ...interface{}) error {
	db := engine.DriveEngine.Table(m.TableName()).Where(condition, param...).Preload("TaxCode").Preload("Order").First(m)
	return db.Error
}

// ReadAllItem to get all data form table item
func ReadAllItem(condition string, param ...interface{}) (*[]Item, int, error) {
	var mx []Item
	db := engine.DriveEngine.Table(new(Item).TableName())
	if condition != "" {
		db = db.Where(condition, param...)
	}
	db = db.Preload("TaxCode").Preload("Order").Find(&mx)

	return &mx, len(mx), db.Error
}
