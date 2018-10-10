package routers

import (
	"tax-calculator/controllers"

	"github.com/labstack/echo"
)

// Handlers interface collect endpoint
type Handlers interface {
	Endpoint(url *echo.Group)
}

// group contain handler func for all module
var group = map[string]Handlers{
	"order":    &controllers.OrderModule{},
	"tax-code": &controllers.TaxcodeModule{},
}

// SetRoute list all group and URL endpoint
func SetRoute(r *echo.Echo) {
	for idx, fc := range group {
		fc.Endpoint(r.Group(idx))
	}
}
