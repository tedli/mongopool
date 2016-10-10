package mongopool

import (
	"testing"
)

func TestMongoPool(t *testing.T) {
	config := defaultConfig{}
	factory := defaultFactory{}
	pool := NewMongoPool(config, factory)
	defer pool.Close()

	rawPool, ok := pool.(*mongoPool)
	if !ok {
		t.Error("not a mongoPool pointer")
	} else {
		t.Log("is a mongoPool pointer as expect")
	}

	initCount := len(rawPool.pool)
	if initCount != config.InitialCount() {
		t.Errorf("initial count is %d not %d specified by config", initCount, config.InitialCount())
	} else {
		t.Logf("initial count is %d as config specified", config.InitialCount())
	}

	sessionA := pool.Get()

	if len(rawPool.pool) != 0 {
		t.Error("not taked from the pool")
	} else {
		t.Log("is taken from the pool")
	}

	sessionB := pool.Get()
	sessionB.Close()

	sessionA.Close()

	if len(rawPool.pool) != 2 {
		t.Error("session count should 2")
	} else {
		t.Log("session count is 2")
	}
	var sessions [11]*MongoSession
	for i := 0; i < config.PoolCapacity(); i++ {
		sessions[i] = pool.Get()
	}
	sessions[10] = pool.Get()

	if len(rawPool.pool) != 0 {
		t.Error("not taked from the pool")
	} else {
		t.Log("is taken from the pool")
	}

	for i := 0; i < 10; i++ {
		sessions[i].Close()
	}

	if len(rawPool.pool) != 10 {
		t.Errorf("pool capacity is not %d", config.PoolCapacity())
	} else {
		t.Logf("pool capacity is %d", config.PoolCapacity())
	}

	if sessions[0].Session != nil {
		t.Error("should be nil")
	} else {
		t.Log("base native session is nil after close so error on Die()")
	}

	sessionC := pool.Get()
	sessions[10].Die()

	if len(rawPool.pool) != 9 {
		t.Error("not close directly")
	} else {
		t.Log("Die() close native session directly")
	}

	sessionC.Close()

	if len(rawPool.pool) != 10 {
		t.Error("not put to the pool")
	} else {
		t.Log("put the wrapped session to the pool")
	}
}
