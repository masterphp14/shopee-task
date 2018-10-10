package controllers

import (
	"tax-calculator/model"

	"github.com/labstack/echo"
)

// TaxcodeModule handling tax_code module
type TaxcodeModule struct{}

// Endpoint list endpoint for all method
func (ctr *TaxcodeModule) Endpoint(url *echo.Group) {
	// no middleware, there is no auth yet..
	url.GET("", ctr.readAll)
	url.GET("/:id", ctr.readOne)
}

// readAll to get all data of tax_code
func (ctr *TaxcodeModule) readAll(ctx echo.Context) error {
	resp := new(model.ResponseFormat)
	mx, total, e := model.ReadAllTaxCode("")
	if e == nil {
		return resp.Serve(ctx, &mx, total)
	}

	return resp.Error(ctx, e)
}

// readOne to get a single data from tax_code table, (id suppose to be encrypt)
func (ctr *TaxcodeModule) readOne(ctx echo.Context) error {
	var e error
	var id int64
	resp := new(model.ResponseFormat)
	if id, e = model.ConvertID(ctx.Param("id")); e == nil {
		m := new(model.Taxcode)
		if e = m.Read("id = ?", id); e == nil {
			return resp.Serve(ctx, m, 1)
		}
	}

	return resp.Error(ctx, e)
}
