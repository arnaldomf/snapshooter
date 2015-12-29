package main

import (
	"github.com/bitly/go-simplejson"
)

const (
	DIGITAL_OCEAN string = "digital_ocean"
	EC2           string = "ec2"
)

type Instance interface {
	Snapshot() bool
}

type Snapshooter struct {
	config    *simplejson.Json
	instances []Instance
}

func (s *Snapshooter) Start() []Instance {
	if _, ok := s.config.CheckGet(DIGITAL_OCEAN); ok {
		do := DigitalOcean{config: s.config.Get(DIGITAL_OCEAN)}

		s.instances = append(s.instances, do.GetInstances()...)
	}

	return s.instances
}
