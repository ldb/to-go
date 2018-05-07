package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Task struct {
	Id          string
	Description string
	Date        time.Time
	Progress    int
	Finished    bool
}

func GetAllTasks() (err error, tasks []Task) {
	results := []Task{}

	err = session.DB(database).C(collection).Find(bson.M{}).All(&results)
	if err != nil {
		return err, []Task{}
	}

	return nil, results
}
