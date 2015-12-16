package models

import "errors"

// Instance is an interface to cloud Instances
type Instance interface {
	windowHour(hh int)
}

// SetWindowHour validates the window_hour attribute and parses it to int
func SetWindowHour(instance Instance, hour int) error {
	if hour < 0 || hour > 23 {
		return errors.New("WindowHour is not valid")
	}
	instance.windowHour(hour)

	return nil
}

// GetInstanceInput packs all input information needed to create an abstract instance
type GetInstanceInput struct {
	Region     string
	WindowHour int
}
