package mqtt

type CLI interface {
	// emqx Config methods
	ConfigSelect(name string) error
	ConfigInfo(name string) (string, error)
	ConfigList() ([]string, error)
	// emqx Cluster methods
	ClusterCheckConnection() (string, error)
	ClusterCheckHealth() (string, error)
	ClusterTopics() ([]string, error)
	ClusterInfo() (string, error)
	// emqx Benchmark methods
	Benchmark() error
}
