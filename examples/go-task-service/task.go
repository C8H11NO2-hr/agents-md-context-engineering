package task

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("task not found")

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type Store struct {
	mu     sync.Mutex
	nextID int
	tasks  map[int]Task
}

func NewStore() *Store {
	return &Store{
		nextID: 1,
		tasks:  make(map[int]Task),
	}
}

func (s *Store) Create(title string) Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := Task{ID: s.nextID, Title: title}
	s.tasks[task.ID] = task
	s.nextID++
	return task
}

func (s *Store) Complete(id int) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return Task{}, ErrNotFound
	}
	task.Completed = true
	s.tasks[id] = task
	return task, nil
}

func (s *Store) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks := make([]Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}
