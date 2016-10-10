package mongopool

import "os"

type defaultConfig struct {
}

func (config defaultConfig) ConnectionString() string {
	return os.Getenv("MONGODB")
}

func (config defaultConfig) InitialCount() int {
	return 1
}

func (config defaultConfig) PoolCapacity() int {
	return 10
}
