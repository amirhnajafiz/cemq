package config

func init() {
	setupBaseConfigDirectories()
}

type Manager interface {
	Select(name string) string
	List() []string
	Info() string
}

func New() Manager {
	return &config{}
}
