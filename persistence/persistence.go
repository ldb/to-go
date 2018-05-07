package persistence

import (
	"fmt"
	"github.com/gin-gonic/gin/json"
	"io"
)

type Task struct {
	Description string `json:"description"`
	Deadline    string `json:"deadline"`
	Progress    int    `json:"progress"`
	Finished    bool   `json:"finished"`
}

type storage struct {
	tasks map[int]Task `json:"tasks"`
	maxId int          `json:"max"`
	store io.ReadWriteCloser
}

func New(store io.ReadWriteCloser) (s storage) {
	return storage{
		tasks: make(map[int]Task),
		maxId: 0,
		store: store,
	}
}

func (s *storage) Save() (err error) {
	data, err := json.Marshal(s.tasks)
	if err != nil {
		return err
	}
	_, err = s.store.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *storage) Insert(task Task) (id int) {
	s.maxId = s.maxId + 1
	s.tasks[s.maxId] = task
	return s.maxId
}

func (s *storage) Get(id int) (err error, task Task) {
	t, ok := s.tasks[id]
	if !ok {
		return fmt.Errorf("task not found"), Task{}
	}
	return nil, t
}

func (s *storage) GetAll() map[int]Task {
	return s.tasks
}

func (s *storage) FindAndUpdate(id int, task Task) (err error) {
	err, _ = s.Get(id)
	if err != nil {
		return err
	}
	s.tasks[id] = task
	return nil
}

func (s *storage) Delete(id int) {
	delete(s.tasks, id)
}
