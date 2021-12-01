package rabbit

import "fmt"

type Config struct {
	Host     string
	User     string
	Password string
	Queue    string
	Port     string
}

func (c *Config) url() string {
	if c.User == "" {
		return fmt.Sprintf("amqp://%s:%s/", c.Host, c.Port)
	}
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", c.User, c.Password, c.Host, c.Port)
}
