package model

import (
	"tax-calculator/engine"
)

// Taxcode model
type Taxcode struct {
	ID         int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Code       string `gorm:"column:code" json:"code"`
	Type       string `gorm:"column:type" json:"type"`
	Refundable string `gorm:"column:refundable" json:"refundable"`
}

// TableName table name in database
func (m *Taxcode) TableName() string {
	return "tax_code"
}

// Read to get one data in tax_code based on parameter
func (m *Taxcode) Read(condition string, param ...interface{}) error {
	db := engine.DriveEngine.Table(m.TableName()).Where(condition, param...).First(m)
	return db.Error
}

// Save to insert data to tax_code table
func (m *Taxcode) Save() error {
	db := engine.DriveEngine.Table(m.TableName()).Create(&m)
	return db.Error
}

// Delete to delete row from table tax_code
func (m *Taxcode) Delete(condition string, param ...interface{}) error {
	db := engine.DriveEngine.Table(m.TableName()).Where(condition, param...).Delete(m)
	return db.Error
}

// ReadAllTaxCode to get all data form table tax_code
func ReadAllTaxCode(condition string, param ...interface{}) (*[]Taxcode, int, error) {
	var mx []Taxcode
	db := engine.DriveEngine.Table(new(Taxcode).TableName())
	if condition != "" {
		db = db.Where(condition, param...)
	}
	db = db.Find(&mx)

	return &mx, len(mx), db.Error
}
