package medium

import (
	medium "github.com/medium/medium-sdk-go"
	"os"
	"errors"
)

const(
	mediumToken = "MEDIUM_ACCESS_TOKEN"
)

type Config struct {
	Client *medium.Medium
	UserID string
}

func (c *Config) LoadAndValidate() error {
	token := os.Getenv(mediumToken)
  if len(token) == 0 {
    return errors.New("define MEDIUM_ACCESS_TOKEN environment variable")
  }
	c.Client = medium.NewClientWithAccessToken(token)
	user, err := c.Client.GetUser("")
	if err != nil {
		return err
	}
	c.UserID = user.ID
	return nil
}