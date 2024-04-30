package model

type Config struct {
	Server      string `json:"server"`
	Description string `json:"description"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
