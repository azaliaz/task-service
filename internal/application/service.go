package application

import (
	"context"
	"fmt"
	"github.com/azaliaz/task-service/internal/storage"
	"sync"
	"time"

	"log/slog"
)

const (
	successResult = "Task completed successfully"
)

type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusRunning   TaskStatus = "running"
	StatusCompleted TaskStatus = "completed"
	StatusDeleted   TaskStatus = "deleted"
)

type Task struct {
	ID          string
	CreatedAt   time.Time
	StartedAt   time.Time
	CompletedAt time.Time
	Status      TaskStatus
	Result      string
}

type TaskService interface {
	CreateTask(ctx context.Context) (string, error)
	DeleteTask(ctx context.Context, id string) error
	GetTaskStatus(ctx context.Context, id string) (*Task, error)
}

type Service struct {
	log     *slog.Logger
	config  *Config
	storage storage.Service
	todo    todo
}

type todo struct {
	mu    sync.RWMutex
	tasks []string
}

func (t *todo) Put(id string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tasks = append(t.tasks, id)
}

func (t *todo) Pull() string {
	t.mu.Lock()
	defer t.mu.Unlock()

	if len(t.tasks) == 0 {
		return ""
	}

	taskId := t.tasks[0]

	t.tasks = t.tasks[1:]

	return taskId
}

func NewService(logger *slog.Logger, config *Config) *Service {
	return &Service{
		log:     logger,
		config:  config,
		storage: storage.NewService(),
	}
}

func (s *Service) Init() error {

	return nil
}

func (s *Service) Run(ctx context.Context) {
	for range 100 {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:

				}

				id := s.todo.Pull()
				if id == "" {
					continue
				}

				err := s.storage.UpdateTaskStatus(id, storage.StatusRunning)
				if err != nil {
					s.log.Error(fmt.Sprintf("failed to update task status: %s", err))
				}

				err = s.storage.UpdateTaskStartedAt(id, time.Now())
				if err != nil {
					s.log.Error(fmt.Sprintf("failed to update task started_at: %s", err))
				}

				s.someWork()

				err = s.storage.UpdateTaskStatus(id, storage.StatusCompleted)
				if err != nil {
					s.log.Error(fmt.Sprintf("failed to update task status: %s", err))
				}

				err = s.storage.UpdateTaskCompletedAt(id, time.Now())
				if err != nil {
					s.log.Error(fmt.Sprintf("failed to update task completed_at: %s", err))
				}

				err = s.storage.UpdateTaskResult(id, successResult)
				if err != nil {
					s.log.Error(fmt.Sprintf("failed to update task result: %s", err))
				}

				s.log.Info("task completed", "id", id)
			}
		}()
	}
}

func (s *Service) Stop() {

}

func (s *Service) someWork() {
	time.Sleep(s.config.TaskDuration)
}
