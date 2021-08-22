package drivers

import (
	"fmt"
	"github.com/polivera/proxy-service/src/data/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteDriver struct {
	db *gorm.DB
}

// NewSQLiteDriver - Return new instance of sqlite driver
func NewSQLiteDriver(path string) (*sqliteDriver, error) {
	var (
		err error
		con *gorm.DB
	)

	if con, err = gorm.Open(sqlite.Open(path), &gorm.Config{}); err != nil {
		fmt.Println("Cannot connect to database")
		return nil, err
	}

	return &sqliteDriver{db: con}, nil
}

func (slt *sqliteDriver) Migrate() error {
	var err error
	if err = slt.db.AutoMigrate(&models.ProxyRequest{}); err != nil {
		return err
	}
	if err = slt.db.AutoMigrate(&models.RequestConfig{}); err != nil {
		return err
	}
	return nil
}
