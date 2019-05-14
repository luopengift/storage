package storage

import (
	"github.com/luopengift/log"
	"gopkg.in/mgo.v2"
)

// MongoDBConfig config
type MongoDBConfig struct {
	URL string `yaml:"url"`
	DB  string `yaml:"db"`
}

var collection *mgo.Database

type Mongo struct {
	*mgo.Database
}

// MongoInit mongo init
func MongoInit(url, db string) (*Mongo, error) {
	mgo.SetDebug(true)
	sess, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Mongo{sess.DB(db)}, nil
}

// Upsert upsert
func Upsert(collect string, index map[string]interface{}, doc interface{}) (*mgo.ChangeInfo, error) {
	info, err := collection.C(collect).Upsert(index, doc)
	return info, err
}

// Query query
func Query(collect string, selector interface{}, results interface{}) error {
	err := collection.C(collect).Find(selector).Sort("-_id").All(results)
	return err
}

// QueryByPage query page
func QueryByPage(collect string, selector, results interface{}, offset, size int) error {
	return collection.C(collect).Find(selector).Sort("-_id").Skip(offset).Limit(size).All(results)
}

// Delete delete
func Delete(collect string, selector interface{}) (*mgo.ChangeInfo, error) {
	info, err := collection.C(collect).RemoveAll(selector)
	log.Debugf("Mongo delete[%s], selector=%#v; info=%#v, err=%v", collect, selector, info, err)
	return info, err
}

// DeleteID delete id
func DeleteID(collect string, selector interface{}) error {
	err := collection.C(collect).RemoveId(selector)
	log.Debugf("Mongo deleteId[%s], selector=%#v; err=%v", collect, selector, err)
	return err
}
