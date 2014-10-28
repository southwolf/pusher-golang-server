package Pusher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var httpClient = http.Client{}

type Client struct {
	appid, app_key, app_secret string
	Host                       string
	Port                       int
}

func CreateClient(appid, key, secret string) *Client {
	return &Client{
		appid:      appid,
		app_key:    key,
		app_secret: secret,
		Host:       "rest.pusher.com",
		Port:       8080,
	}
}

func (c *Client) Trigger(data, event string, channels ...string) error {

}

func (c *Client) TriggerPath() string {

}

func (c *Client) send(absolutePath string, queryString string, content []byte) {

}
