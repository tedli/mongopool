package mongopool

import "gopkg.in/mgo.v2"

type Config interface {
	ConnectionString() string
	InitialCount() int
	PoolCapacity() int
}

type Factory interface {
	CreateNativeSession(string) *mgo.Session
}

type MongoPool interface {
	Get() *MongoSession
	Close()
}

func NewMongoPool(config Config, factory Factory) MongoPool {
	return newMongoPool(config, factory)
}