package service

import (
	"github.com/layroscloud/todo-go/entity"
	"github.com/layroscloud/todo-go/pkg/dto"
	"github.com/layroscloud/todo-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username string, password string) (string, error)
	ParseToken(accessToken string) (int64, error)
}

type TodoList interface {
	FindAllLists(userId int64) ([]entity.TodoList, error)
	FindByIdList(id int64) (*entity.TodoList, error)
	CreateList(dto dto.ListCreateDto, userId int64) (int64, error)
	UpdateList(dto dto.ListUpdateDto) (int64, error)
	DeleteList(id int64) (int, error)
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		TodoList:      NewListService(repository.TodoList),
	}
}
