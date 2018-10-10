package engine

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DriveEngine database connection (make it global)
var DriveEngine *gorm.DB

// Start to open connection to database (for driver and time parse can be change)
func Start(connection string) {
	var err error
	// open connection ("mysql", "user:password@tcp(host)/database")
	DriveEngine, err = gorm.Open("mysql", connection+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	} else {
		// set gorm table name to singular and allow log mode to view query
		DriveEngine.SingularTable(true)
		DriveEngine.LogMode(true)
	}
}

// Stop to close connection to database
func Stop() error {
	return DriveEngine.Close()
}
