package main

import (
	"github.com/bitly/go-simplejson"
	api "github.com/digitalocean/godo"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
	"testing"
)

func TestZeroInstancesIfDropletsNotExists(t *testing.T) {
	config, _ := simplejson.NewJson([]byte(`{}`))

	do := &DigitalOcean{config: config}

	assert := assert.New(t)
	assert.Zero(do.GetInstances())
}

func TestGetInstances(t *testing.T) {
	config, _ := simplejson.NewJson([]byte(`{
		"droplets": [{"id": 1, "name": "Goku", "window_hour": 10 }]
	}`))

	do := &DigitalOcean{config: config}

	instances := do.GetInstances()

	assert := assert.New(t)
	assert.Len(instances, 1)
	assert.Equal(instances[0], &DigitalOceanInstance{id: 1, name: "Goku", windowHour: 10})
}

func TestSnapshot(t *testing.T) {
	doInstance := new(DigitalOceanInstance)

	assert := assert.New(t)
	assert.Equal(doInstance.Snapshot(), false)
}

func TestGenerateOauthToken(t *testing.T) {
	tokenSource := &TokenSource{
		AccessToken: "AAA",
	}

	token, _ := tokenSource.Token()

	assert := assert.New(t)
	assert.IsType(&oauth2.Token{}, token)
}

func TestClient(t *testing.T) {
	config, _ := simplejson.NewJson([]byte(`{
		"access_token": "AAA"
	}`))

	do := &DigitalOcean{config: config}

	assert := assert.New(t)
	assert.IsType(&api.Client{}, do.Client())

	// Pre loaded
	assert.IsType(&api.Client{}, do.Client())

}
