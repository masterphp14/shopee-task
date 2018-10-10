package controllers

import (
	"tax-calculator/model"

	"fmt"

	"github.com/labstack/echo"
)

// OrderModule handling order module
type OrderModule struct{}

// Endpoint list endpoint for all method
func (ctr *OrderModule) Endpoint(url *echo.Group) {
	// no middleware, there is no auth yet..
	url.GET("", ctr.readAll)
	url.GET("/:id", ctr.readOne)
	url.POST("", ctr.create)
	//url.PUT("/:id", ctr.update)
	//url.DELETE("/:id", ctr.delete)
}

// readOne to get a single data from tax_code table, (id suppose to be encrypt)
func (ctr *OrderModule) readOne(ctx echo.Context) error {
	var e error
	var id int64
	resp := new(model.ResponseFormat)
	if id, e = model.ConvertID(ctx.Param("id")); e == nil {
		m := new(model.Order)
		if e = m.Read("id = ?", id); e == nil {
			m.Calculate()
			return resp.Serve(ctx, m, 1)
		}
	}

	return resp.Error(ctx, e)
}

// readAll to get all data of order
func (ctr *OrderModule) readAll(ctx echo.Context) error {
	resp := new(model.ResponseFormat)
	mx, total, e := model.ReadAllOrder("")
	if e == nil {
		return resp.Serve(ctx, &mx, total)
	}

	return resp.Error(ctx, e)
}

// create to insert data to order
func (ctr *OrderModule) create(ctx echo.Context) error {
	resp := new(model.ResponseFormat)
	var e error
	m := new(model.Order)
	if e = ctx.Bind(m); e == nil {
		if resp.Message(m.ValidateCreate()) {
			fmt.Println("aaaa")
			if e = m.Create(); e == nil {
				if e = m.Read("id = ?", m.ID); e == nil {
					m.Calculate()
					return resp.Serve(ctx, m, 1)
				}
			}
		}
	}
	fmt.Println("bbbb")
	return resp.Error(ctx, e)
}
