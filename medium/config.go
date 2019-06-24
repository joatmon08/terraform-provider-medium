package medium

import (
	medium "github.com/medium/medium-sdk-go"
)

type Config struct {
	AccessToken string
	Client *medium.Medium
}

func (c *Config) LoadAndValidate() {
	c.Client = medium.NewClientWithAccessToken(c.AccessToken)
}