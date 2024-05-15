package model

// Config struct is a holder for each conf.json file in config/emqx directory
type Config struct {
	Server      string `json:"server"`
	Description string `json:"description"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
