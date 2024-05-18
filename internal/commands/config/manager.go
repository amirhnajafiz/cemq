package config

func init() {
	setupBaseConfigDirectories()
}

// Manager handles the configs CRUD operations
type Manager interface {
	Select(name string) string
	List() []string
	Info() string
}

func New() Manager {
	return &config{}
}
