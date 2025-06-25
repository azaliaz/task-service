package application

import (
	"context"
	"fmt"
)

func (s *Service) CreateTask(ctx context.Context) (string, error) {
	task, err := s.storage.CreateTask()
	if err != nil {
		return "", fmt.Errorf("failed to create task: %w", err)
	}

	s.todo.Put(task.ID)

	s.log.Info("task created", "id", task.ID)

	return task.ID, nil
}

func (s *Service) DeleteTask(ctx context.Context, id string) error {
	err := s.storage.DeleteTask(id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	s.log.Info("task deleted", "id", id)

	return nil
}

func (s *Service) GetTaskStatus(ctx context.Context, id string) (*Task, error) {
	task, err := s.storage.GetTaskStatus(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task status: %w", err)
	}

	return &Task{
		ID:          task.ID,
		CreatedAt:   task.CreatedAt,
		StartedAt:   task.StartedAt,
		CompletedAt: task.CompletedAt,
		Status:      TaskStatus(task.Status),
		Result:      task.Result,
	}, nil
}
