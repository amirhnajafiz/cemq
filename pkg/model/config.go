package model

import "fmt"

// Config struct is a holder for each conf.json file in config/emqx directory
type Config struct {
	Server      string `json:"server"`
	Port        int    `json:"port"`
	Description string `json:"description"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

func (c Config) URL() string {
	return fmt.Sprintf("tcp://%s:%d", c.Server, c.Port)
}
