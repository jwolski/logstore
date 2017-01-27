package main

import (
	"errors"

	"gopkg.in/mgo.v2"
)

var (
	errDialing = errors.New("failed to dial")
	errSaving  = errors.New("failed to save")
)

// Stores log records to persistence storage.
type store interface {
	close()
	save(*logRecord) error
}

// Configures MongoDB
type mongoConfig struct {
	url        string
	db         string
	collection string
}

var defaultMongoConf = mongoConfig{
	url:        "localhost",
	db:         "logstore",
	collection: "logs",
}

// Encapsulates access to MongoDB
type mongoStore struct {
	conf    mongoConfig
	session *mgo.Session
}

// Creates a MongoDB wrapper that encapsulates data access to MongoDB
func dialMongo(conf mongoConfig) (store, error) {
	mongo := &mongoStore{conf: conf}
	// TODO: Understand session management within mgo.
	session, err := mgo.Dial(conf.url)
	if err != nil {
		return mongo, err
	}
	mongo.session = session
	return mongo, nil
}

// Closes outstanding Mongo session
func (s *mongoStore) close() {
	if s.session == nil {
		return
	}
	s.session.Close()
}

// Saves the log record
func (s *mongoStore) save(r *logRecord) error {
	coll := s.session.DB(s.conf.db).C(s.conf.collection)
	if err := coll.Insert(r); err != nil {
		return err
	}
	return nil
}
