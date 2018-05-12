package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Task struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Description string        `json:"description" bson:"description"`
	Date        time.Time     `json:"date" bson:"date"`
	Progress    int           `json:"progress" bson:"progress"`
	Finished    bool          `json:"finished" bson:"finished"`
}

func (s *server) GetAllTasks() (err error, tasks []Task) {
	results := []Task{}

	err = s.C.Find(bson.M{}).All(&results)
	if err != nil {
		return err, []Task{}
	}

	return nil, results
}

func (s *server) FindTask(id string) (err error, task Task) {
	result := Task{}

	err = s.C.Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&result)
	if err != nil {
		return err, Task{}
	}

	return nil, result
}

func (s *server) InsertTask(t Task) (err error, task Task) {
	i := bson.NewObjectId()
	t.Id = i

	err = s.C.Insert(t)
	if err != nil {
		return err, Task{}
	}
	return nil, t
}

func (s *server) UpdateTask(id string, t Task) (err error, task Task) {
	err = s.C.UpdateId(id, task)
	if err != nil {
		return err, Task{}
	}
	return nil, t
}

func (s *server) DeleteTask(id string) (err error) {
	return s.C.RemoveId(id)
}
