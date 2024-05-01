package config

type Manager interface {
	Select(name string) string
	List() []string
	Info() string
}
