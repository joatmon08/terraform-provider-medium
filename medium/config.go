package medium

import (
	"errors"
	"os"

	medium "github.com/medium/medium-sdk-go"
)

const (
	mediumToken = "MEDIUM_ACCESS_TOKEN"
)

type Config struct {
	Client       *medium.Medium
	User         *medium.User
	ReadEndpoint *ReadEndpoint
}

func (c *Config) LoadAndValidate() error {
	token := os.Getenv(mediumToken)
	if len(token) == 0 {
		return errors.New("define MEDIUM_ACCESS_TOKEN environment variable")
	}
	c.ReadEndpoint = &ReadEndpoint{
		Host:        ReadEndpointHost,
		AccessToken: token,
	}

	c.Client = medium.NewClientWithAccessToken(token)
	user, err := c.Client.GetUser("")
	if err != nil {
		return err
	}
	c.User = user
	return nil
}
