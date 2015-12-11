package models

import (
	"errors"
	"strconv"
)

// Instance is an interface to cloud Instances
type Instance interface {
	windowHour(hh int)
}

// SetWindowHour validates the window_hour attribute and parses it to int
func SetWindowHour(instance Instance, hour string) error {
	hh, err := strconv.Atoi(hour)
	if err != nil {
		return err
	}
	if hh < 0 || hh > 23 {
		return errors.New("WindowHour is not valid")
	}
	instance.windowHour(hh)

	return nil
}
