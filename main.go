package main

import (
	"tax-calculator/setup"
)

func main() {
	setup.Initialize()
	// address port is temp
	setup.Run(":1234")

	defer setup.Close()
}
