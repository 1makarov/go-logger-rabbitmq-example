package postgres

import "fmt"

type Config struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func (c *Config) string() string {
	return fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s",
		c.Host, c.User, c.Password, c.Port, c.Name)
}
