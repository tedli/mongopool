package mongopool

import "os"

type DefaultConfig struct {
}

func (config DefaultConfig) ConnectionString() string {
	return os.Getenv("MONGODB")
}

func (config DefaultConfig) InitialCount() int {
	return 1
}

func (config DefaultConfig) PoolCapacity() int {
	return 10
}
