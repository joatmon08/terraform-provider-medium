package medium

import (
	"errors"
	"os"

	read "github.com/joatmon08/terraform-provider-medium/readmedium"
	medium "github.com/medium/medium-sdk-go"
)

const (
	mediumToken = "MEDIUM_ACCESS_TOKEN"
)

type Config struct {
	Client        *medium.Medium
	User          *medium.User
	StoryEndpoint *read.StoryEndpoint
}

func (c *Config) LoadAndValidate() error {
	token := os.Getenv(mediumToken)
	if len(token) == 0 {
		return errors.New("define MEDIUM_ACCESS_TOKEN environment variable")
	}
	c.StoryEndpoint = &read.StoryEndpoint{
		Host: read.MediumURL,
	}

	c.Client = medium.NewClientWithAccessToken(token)
	user, err := c.Client.GetUser("")
	if err != nil {
		return err
	}
	c.User = user
	return nil
}
