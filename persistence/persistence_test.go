package persistence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

}

func TestStorage_Insert(t *testing.T) {
	task := Task{
		Description: "Desc",
		Deadline:    "2018-05-14",
		Progress:    0,
		Finished:    false,
	}

	s := storage{
		tasks: make(map[int]Task),
	}
	id := s.Insert(task)

	assert.Equal(t, 1, id, "first id should be 1")
	assert.Contains(t, s.tasks, id, "task should be saved at correct id")
	assert.Equal(t, s.tasks[id], task, "task should be saved")
}

func TestStorage_Get(t *testing.T) {
	task := Task{
		Description: "Desc1",
		Deadline:    "2018-05-14",
		Progress:    0,
		Finished:    false,
	}

	s := storage{
		tasks: map[int]Task{
			1: task,
		},
	}

	err, taskResult := s.Get(1)

	assert.Nil(t, err, "no error")
	assert.Equal(t, taskResult, task, "task should be correct")

	err, taskResult = s.Get(2)
	assert.EqualError(t, err, "task not found", "should return correct error if task is not found")
}

func TestStorage_Delete(t *testing.T) {
	task := Task{
		Description: "Desc1",
		Deadline:    "2018-05-14",
		Progress:    0,
		Finished:    false,
	}

	s := storage{
		tasks: map[int]Task{
			1: task,
		},
	}

	s.Delete(1)
	assert.Equal(t, s.tasks, make(map[int]Task), "task should be deleted")
}

func TestStorage_GetAll(t *testing.T) {
	task := Task{
		Description: "Desc1",
		Deadline:    "2018-05-14",
		Progress:    0,
		Finished:    false,
	}

	s := storage{
		tasks: map[int]Task{
			1: task,
			2: task,
		},
	}

	l := s.GetAll()

	assert.Equal(t, l, s.tasks, "should return all tasks")
}

func TestStorage_FindAndUpdate(t *testing.T) {
	task := Task{
		Description: "Desc1",
		Deadline:    "2018-05-14",
		Progress:    0,
		Finished:    false,
	}

	taskUpdate := Task{
		Description: "Desc2",
		Deadline:    "2018-05-14",
		Progress:    10,
		Finished:    false,
	}

	s := storage{
		tasks: map[int]Task{
			1: task,
			2: task,
		},
	}

	err := s.FindAndUpdate(3, taskUpdate)
	assert.EqualError(t, err, "task not found", "should return correct error if task is not found")

	err = s.FindAndUpdate(1, taskUpdate)
	assert.Nil(t, err, "should not result in error")
	assert.Equal(t, s.tasks[1], taskUpdate, "should correctly update")
	assert.Equal(t, s.tasks[2], task, "should not update wrong task")
}
