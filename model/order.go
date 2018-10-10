package model

import (
	"fmt"
	"tax-calculator/engine"
	"time"
)

// Order model
type Order struct {
	ID            int       `gorm:"primary_key;AUTO_INCREMENT" json:"id,omitempty"`
	NoOrder       string    `gorm:"column:no_order" json:"no_order"`
	Date          time.Time `gorm:"column:date" json:"date"`
	Items         []Item    `gorm:"foreignkey:OrderID" json:"items,omitempty"`
	PriceSubtotal float64   `gorm:"-" json:"price_subtotal,omitempty"`
	TaxSubtotal   float64   `gorm:"-" json:"tax_subtotal,omitempty"`
	GrandSubtotal float64   `gorm:"-" json:"grand_subtotal,omitempty"`
}

// TableName table name in database
func (m *Order) TableName() string {
	return "order"
}

// Read to get one data in order based on parameter (using preload to get foreign key; not efficient)
func (m *Order) Read(condition string, param ...interface{}) error {
	db := engine.DriveEngine.Table(m.TableName()).Where(condition, param...).Preload("Items").Preload("Items.Order").Preload("Items.TaxCode").First(m)
	return db.Error
}

// Create to insert data to table order, etc
func (m *Order) Create() error {
	db := engine.DriveEngine.Begin()
	defer func() {
		if e := recover(); e != nil {
			db.Rollback()
		}
	}()
	if e := db.Error; e != nil {
		return e
	}
	mx := m.Items
	m.Items = []Item{}
	if e := db.Create(m).Error; e != nil {
		db.Rollback()
		return e
	} else {
		for _, x := range mx {
			x.OrderID = m.ID
			if e = db.Table(new(Item).TableName()).Create(&x).Error; e != nil {
				db.Rollback()
				return e
			}
			m.Items = append(m.Items, x)
		}
	}

	return db.Commit().Error
}

// Calculate to calculate tax and amount
func (m *Order) Calculate() {
	var totalAmount, taxAmount, subPrice float64
	mx := m.Items
	m.Items = []Item{}
	// i supposed to use a pointer here, but i'm in hurry so..
	for _, x := range mx {
		subPrice += x.Price
		//calc tax
		if x.TaxCodeID == int(1) {
			x.Tax = (10 * x.Price) / 100
		} else if x.TaxCodeID == int(2) {
			x.Tax = 10 + ((2 * x.Price) / 100)
		} else if x.TaxCodeID == int(3) {
			if x.Price > float64(100) {
				x.Tax = (1 * (x.Price - float64(100))) / 100
			}
		}
		//....
		x.Amount = x.Price + x.Tax

		taxAmount += x.Tax
		totalAmount += x.Amount
		m.Items = append(m.Items, x)
	}
	m.TaxSubtotal = taxAmount
	m.PriceSubtotal = subPrice
	m.GrandSubtotal = totalAmount
}

// validate to check order data is valid or not
func (m *Order) ValidateCreate() (e map[string]string) {
	e = make(map[string]string)
	// validate if item empty or not
	if len(m.Items) == int(0) {
		e["items"] = "cannot be empty"
	} else {
		for idx, x := range m.Items {
			// validate item name, tax_code_id
			if x.TaxCodeID == int(0) {
				e[fmt.Sprintf("tax_code_id.%d", idx)] = "cannot be 0"
			} else {
				// validate if tax code is exist in database
				tx := new(Taxcode)
				if err := tx.Read("id = ?", x.TaxCodeID); err != nil {
					e[fmt.Sprintf("tax_code_id.%d", idx)] = err.Error()
				}
			}
			if x.ItemName == "" {
				e[fmt.Sprintf("item_name.%d", idx)] = "cannot be empty"
			}
		}
	}
	return
}

// ReadAllOrder to get all data form table tax_code
func ReadAllOrder(condition string, param ...interface{}) (*[]Order, int, error) {
	var mx []Order
	db := engine.DriveEngine.Table(new(Order).TableName())
	if condition != "" {
		db = db.Where(condition, param...)
	}
	db = db.Find(&mx)

	return &mx, len(mx), db.Error
}
