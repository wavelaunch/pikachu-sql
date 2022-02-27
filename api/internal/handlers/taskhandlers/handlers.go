package taskhandlers

import "github.com/wavelaunch/pikachu-sql/api/internal/ports"

type handler struct {
	taskService ports.TaskService
}

func New(taskService ports.TaskService) *handler {
	return &handler{
		taskService: taskService,
	}
}
