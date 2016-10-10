# MongoDB session pool in golang

It's a tiny simple implementation of a pool holding native mongodb session provided by [mongodb golang driver](http://labix.org/mgo). 


## How to use

Just provide the `config` and the `factory`. You can hard code the connection string in your factory, and return what ever in config since you just ignore the parameter passed into the factory.


### Give a config

As the name. It provide the configuration.

```
type Config interface {
	ConnectionString() string
	InitialCount() int
	PoolCapacity() int
}
```


### Give a factory

It provide the native mongodb session.

```
type Factory interface {
	CreateNativeSession(string) *mgo.Session
}
```


### Get a pool

```
pool := NewMongoPool(config, factory)
defer pool.Close()
```

The `pool.Close()` close the pool and all the session it held.


### Get a session from a pool

```
session := pool.Get()
```

Then you can use the warpped session instance as if it is a native session.


### Put a session (taken from a pool) back to the pool it belongs to

```
session.Close()
```

You can just close it. Then it will be put back to the pool.


### Close the backend native session for some reason

When you are using a session taken from a pool, and the session may be dead for some reason, so it's better not to put it back. You can close the backend native session like this.

```
session.Die()
```

**Notice:** When a warpped session is closed, then the backend native session object is set to `nil`. So you will get panic when you die a closed session.
