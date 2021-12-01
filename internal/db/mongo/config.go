package mongo

import "fmt"

type Config struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func (c *Config) url() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/", c.User, c.Password, c.Host, c.Port)
}
