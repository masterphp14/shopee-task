// Copyright 2016 PT. Qasico Teknologi Indonesia. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package test

import (
	"tax-calculator/engine"
	"tax-calculator/model"
	"tax-calculator/setup"
)

// Setup testing setup.
func Setup() {
	setup.Initialize()
	// address port is temp
	setup.Run(":1234")

	defer setup.Close()
}

// DbClean cleaning all data from table order and item.
func DbClean() {
	engine.DriveEngine.Delete(&model.Order{})
}
