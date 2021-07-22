package data

import "errors"

type Store interface {

}

func NewStore(driver string) (Store, error) {
	switch driver {
	default:
		return nil, errors.New("Driver " + driver + " does not exist")
	}
}
