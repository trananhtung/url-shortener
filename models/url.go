package models

import (
	"fmt"
	"os"
)

type URL struct {
	PROTOCOL string
	HOST     string
	PORT     string
}

func (c *URL) Init() {
	fmt.Println("Initializing config...")

	c.PROTOCOL = os.Getenv("PROTOCOL")
	c.HOST = os.Getenv("HOST")
	c.PORT = os.Getenv("PORT")
}

func (c *URL) GetURL() string {
	return fmt.Sprintf("%s://%s:%s", c.PROTOCOL, c.HOST, c.PORT)
}
