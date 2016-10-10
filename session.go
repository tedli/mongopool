package mongopool

import "gopkg.in/mgo.v2"

type MongoSession struct {
	*mgo.Session
	pool *mongoPool
}

func (ms *MongoSession) Close() {
	ms.pool.put(ms.Session)
	ms.Session = nil
}

func (ms *MongoSession) Die() {
	ms.Session.Close()
}