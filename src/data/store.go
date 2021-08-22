package data

import (
	"errors"
	"github.com/polivera/proxy-service/src/data/drivers"
)

type Store interface {
	Migrate() error
}

func NewStore(driver string, path string) (Store, error) {
	switch driver {
	case "sqlite":
		return drivers.NewSQLiteDriver(path)
	default:
		return nil, errors.New("Driver " + driver + " does not exist")
	}
}
