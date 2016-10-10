package mongopool

import (
	"gopkg.in/mgo.v2"
	"sync"
)

type mongoPool struct {
	config Config
	factory Factory
	syncRoot sync.Mutex
	pool chan *mgo.Session
}

func (mp *mongoPool) Get() MongoSession {
	mp.syncRoot.Lock()
	defer mp.syncRoot.Unlock()
	select {
	case session := <- mp.pool:
		return MongoSession{session, mp}
	default:
		return MongoSession{mp.factory.CreateNativeSession(mp.config.ConnectionString()), mp}
	}
}

func (mp *mongoPool) Close() {
	mp.syncRoot.Lock()
	defer mp.syncRoot.Unlock()
	count := len(mp.pool)
	for i := 0; i < count; i++ {
		session := <- mp.pool
		session.Close()
	}
	close(mp.pool)
}

func newMongoPool(config Config, factory Factory) *mongoPool {
	pool := &mongoPool{
		config: config,
		factory: factory,
		pool: make(chan *mgo.Session, config.PoolCapacity()),
	}
	for i := 0; i < config.InitialCount(); i++ {
		pool.pool <- factory.CreateNativeSession(config.ConnectionString())
	}
	return pool
}

func (mp *mongoPool) put(session *mgo.Session) {
	mp.syncRoot.Lock()
	defer mp.syncRoot.Unlock()
	count := len(mp.pool)
	if count >= mp.config.PoolCapacity() {
		session.Close()
	} else {
		mp.pool <- session
	}
}
