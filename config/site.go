package config

import (
	"fmt"
	"os"
)

type SITE struct {
	PROTOCOL string
	HOST     string
	PORT     string
}

func (c *SITE) Init() {
	fmt.Println("Initializing config...")

	c.PROTOCOL = os.Getenv("PROTOCOL")
	c.HOST = os.Getenv("HOST")
	c.PORT = os.Getenv("PORT")
}

func (c *SITE) GetSite() string {
	return fmt.Sprintf("%s://%s:%s", c.PROTOCOL, c.HOST, c.PORT)
}
