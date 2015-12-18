package main

import (
	j "github.com/bitly/go-simplejson"
	api "github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type DigitalOceanInstance struct {
	id         int
	name       string
	windowHour int
}

func (do *DigitalOceanInstance) Snapshot() bool {
	return false
}

type DigitalOcean struct {
	config *j.Json
	client *api.Client
}

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}

	return token, nil
}

func (do *DigitalOcean) Client() *api.Client {

	if do.client != nil {
		return do.client
	}

	auth_token, _ := do.config.Get("auth_token").String()

	tokenSource := &TokenSource{
		AccessToken: auth_token,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := api.NewClient(oauthClient)

	do.client = client

	return client
}

func (do *DigitalOcean) GetDropletIdByName(name string) int {
	opt := &api.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	droplets, _, err := do.Client().Droplets.List(opt)

	if err != nil {
		return 0
	}

	for _, droplet := range droplets {
		if droplet.Name == name {
			return droplet.ID
		}
	}

	return 0
}

func (do *DigitalOcean) GetInstances() []Instance {
	var instances []Instance

	if _, ok := do.config.CheckGet("droplets"); ok == false {
		return instances
	}

	droplets := do.config.Get("droplets")

	for key, _ := range droplets.MustArray() {

		id := 0

		if value, ok := droplets.GetIndex(key).CheckGet("id"); ok {
			id, _ = value.Int()
		}

		name, _ := droplets.GetIndex(key).Get("name").String()
		windowHour, _ := droplets.GetIndex(key).Get("window_hour").Int()

		if id == 0 {
			id = do.GetDropletIdByName(name)
		}

		if id != 0 {
			instance := &DigitalOceanInstance{id: id, name: name, windowHour: windowHour}
			instances = append(instances, instance)
		}
	}

	return instances
}
