package service

import (
	"github.com/layroscloud/todo-go/entity"
	"github.com/layroscloud/todo-go/pkg/dto"
	"github.com/layroscloud/todo-go/pkg/repository"
)

type ListService struct {
	repo repository.TodoList
}

func NewListService(repo repository.TodoList) *ListService {
	return &ListService{
		repo: repo,
	}
}

func (l *ListService) FindAllLists(userId int64) ([]entity.TodoList, error) {
	items, err := l.repo.FindAllByUserId(userId)
	return items, err
}

func (l *ListService) FindByIdList(id int64) (*entity.TodoList, error) {
	item, err := l.repo.FindById(id)
	return item, err
}

func (l *ListService) CreateList(dto dto.ListCreateDto, userId int64) (int64, error) {
	list := entity.TodoList{
		Title:       dto.Title,
		Description: dto.Description,
	}
	id, err := l.repo.Create(list, userId)
	return id, err
}

func (l *ListService) UpdateList(dto dto.ListUpdateDto) (int64, error) {
	list := entity.TodoList{
		Id:          dto.Id,
		Title:       dto.Title,
		Description: dto.Description,
	}
	index, err := l.repo.Update(list)
	return index, err
}

func (l *ListService) DeleteList(id int64) (int, error) {
	index, err := l.repo.Delete(id)
	return index, err
}
