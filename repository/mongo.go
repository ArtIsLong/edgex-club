package repository

import (
	"edgex-club/config"
	"gopkg.in/mgo.v2"
	"time"
)

type DataStore struct {
	S *mgo.Session
}

var DS DataStore

func (ds DataStore) DataStore() *DataStore {
	return &DataStore{ds.S.Copy()}
}

func DBConnect() bool {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.Conf.DB.Addr},
		Timeout:  time.Duration(5000) * time.Millisecond,
		Database: config.Conf.DB.Database,
		Username: config.Conf.DB.Username,
		Password: config.Conf.DB.Password,
	}
	s, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return false
	}
	s.SetSocketTimeout(time.Duration(5000) * time.Millisecond)
	DS.S = s
	return true
}
