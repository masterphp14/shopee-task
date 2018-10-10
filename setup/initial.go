package setup

import (
	"fmt"
	"os"
	"tax-calculator/engine"
	"tax-calculator/routers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

//Router driver for router
var Router *echo.Echo

func Initialize() {
	conn := LoadEnvironment()
	engine.Start(conn)

	// create new routing and set URL
	Router = echo.New()
	routers.SetRoute(Router)
}

// Close to stop connection to database (use with defer)
func Close() error {
	return engine.Stop()
}

// Run start building server for API
func Run(addr string) {
	if err := Router.Start(addr); err != nil {
		Router.Logger.Fatal(err)
	}
}

// LoadEnvironment read parameter from .env
func LoadEnvironment() string {
	godotenv.Load()
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	db := os.Getenv("DB_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, db)
}
