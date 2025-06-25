package storage

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

func (s *service) CreateTask() (*Task, error) {
	id := uuid.NewString()

	task := &Task{
		ID:        id,
		CreatedAt: time.Now(),
		Status:    StatusPending,
	}

	s.mu.Lock()
	s.tasks[id] = task
	s.mu.Unlock()

	return task, nil
}

func (s *service) DeleteTask(id string) error {
	s.mu.RLock()
	_, exists := s.tasks[id]
	s.mu.RUnlock()
	if !exists {
		return errors.New("task not found")
	}

	s.mu.Lock()
	delete(s.tasks, id)
	s.mu.Unlock()

	return nil
}

func (s *service) GetTaskStatus(id string) (*Task, error) {
	s.mu.RLock()
	task, exists := s.tasks[id]
	s.mu.RUnlock()
	if !exists {
		return nil, errors.New("task not found")
	}

	return task, nil
}

func (s *service) UpdateTaskStatus(id string, status TaskStatus) error {
	s.mu.RLock()
	task, exists := s.tasks[id]
	s.mu.RUnlock()
	if !exists {
		return errors.New("task not found")
	}

	task.Status = status

	s.mu.Lock()
	s.tasks[id] = task
	s.mu.Unlock()

	return nil
}

func (s *service) UpdateTaskStartedAt(id string, startedAt time.Time) error {
	s.mu.RLock()
	task, exists := s.tasks[id]
	s.mu.RUnlock()
	if !exists {
		return errors.New("task not found")
	}

	task.StartedAt = startedAt

	s.mu.Lock()
	s.tasks[id] = task
	s.mu.Unlock()

	return nil
}

func (s *service) UpdateTaskCompletedAt(id string, completedAt time.Time) error {
	s.mu.RLock()
	task, exists := s.tasks[id]
	s.mu.RUnlock()
	if !exists {
		return errors.New("task not found")
	}

	task.CompletedAt = completedAt

	s.mu.Lock()
	s.tasks[id] = task
	s.mu.Unlock()

	return nil
}

func (s *service) UpdateTaskResult(id string, result string) error {
	s.mu.RLock()
	task, exists := s.tasks[id]
	s.mu.RUnlock()
	if !exists {
		return errors.New("task not found")
	}

	task.Result = result

	s.mu.Lock()
	s.tasks[id] = task
	s.mu.Unlock()

	return nil
}
