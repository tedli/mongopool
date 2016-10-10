package mongopool

import "gopkg.in/mgo.v2"

type DefaultFactory struct {
}

func (factory DefaultFactory) CreateNativeSession(connectionString string) *mgo.Session {
	session, err := mgo.Dial(connectionString)
	if err != nil {
		panic(err.Error())
	}
	session.SetMode(mgo.Primary, true)
	return session
}
