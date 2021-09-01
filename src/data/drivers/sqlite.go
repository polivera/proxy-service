package drivers

import (
	"fmt"
	"github.com/polivera/proxy-service/src/data/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
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

	// Query logger, remove or make it configurable
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	if con, err = gorm.Open(sqlite.Open(path), &gorm.Config{Logger: newLogger}); err != nil {
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

func (slt *sqliteDriver) GetConfig(host string) (models.RequestConfig, error) {
	var config models.RequestConfig
	result := slt.db.First(&config, "source = ?", host)
	return config, result.Error
}
