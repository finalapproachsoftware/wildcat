package biz

import mgo "gopkg.in/mgo.v2"

type MongoCollection struct {
	Session        *mgo.Session
	DbName         string
	CollectionName string
}

func NewMongoCollection(s *mgo.Session, dbName string, collectionName string) (error, *MongoCollection) {
	m := new(MongoCollection)
	m.Session = s
	m.DbName = dbName
	m.CollectionName = collectionName

	return nil, m
}

func (c *MongoCollection) Add(docs ...interface{}) error {
	session := c.Session.Copy()
	defer session.Close()

	mongo := session.DB(c.DbName).C(c.CollectionName)
	err := mongo.Insert(docs)

	return err
}
