package main

import (
	"github.com/bitly/go-simplejson"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartReturnZeroInstances(t *testing.T) {
	config, _ := simplejson.NewJson([]byte(`{
	}`))

	snap := &Snapshooter{config: config}

	assert := assert.New(t)
	assert.Zero(snap.Start())
}

func TestStartReturnInstances(t *testing.T) {
	config, _ := simplejson.NewJson([]byte(`{
        "digital_ocean": {"droplets": [{"id": 1, "name": "Goku", "window_hour": 10 }]}
	}`))

	snap := &Snapshooter{config: config}

	assert := assert.New(t)
	assert.Len(snap.Start(), 1)
}
