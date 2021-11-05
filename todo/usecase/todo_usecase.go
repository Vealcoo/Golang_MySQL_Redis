package usecase

import (
	"CleanTodo/model"
	"errors"
	"strconv"
)

type todoUsecase struct {
	todoRepo  model.TodoRepository
	todoCache model.TodoCache
}

func NewTodoUsecase(r model.TodoRepository, c model.TodoCache) model.TodoUsecase {
	return &todoUsecase{
		todoRepo:  r,
		todoCache: c,
	}
}

func (t *todoUsecase) Create(title, content string) (int64, error) {
	if title == "" || content == "" {
		return 0, errors.New("title or content is nil!")
	}
	result, err := t.todoRepo.Create(title, content)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (t *todoUsecase) Delete(listid string) error {
	if listid == "" {
		return errors.New("input have nil")
	}
	_, err := strconv.Atoi(listid)
	if err != nil {
		return errors.New("check input again!")
	}
	//Delete Cache
	if t.todoCache.CacheLoad(listid) {
		t.todoCache.CacheDel(listid)
	}
	result, err := t.todoRepo.Delete(listid)
	if err != nil {
		return err
	}
	if result == 0 {
		return errors.New("can't find listid: " + listid)
	}
	return nil
}

func (t *todoUsecase) Update(listid, title, content string) error {
	if listid == "" || title == "" || content == "" {
		return errors.New("input have nil!")
	}
	_, err := strconv.Atoi(listid)
	if err != nil {
		return errors.New("check input again!")
	}
	//Update Cache
	if t.todoCache.CacheLoad(listid) {
		t.todoCache.CacheSet(listid, title, content)
	}
	result, err := t.todoRepo.Update(listid, title, content)
	if result == 0 {
		return errors.New("can't find listid: " + listid)
	}
	return nil
}

func (t *todoUsecase) GetAll() ([]model.List, error) {
	var result []model.List
	result, err := t.todoRepo.GetAll()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (t *todoUsecase) GetOne(listid string) (model.List, error) {
	if listid == "" {
		return model.List{}, errors.New("input have nil")
	}
	_, err := strconv.Atoi(listid)
	if err != nil {
		return model.List{}, errors.New("check input again")
	}
	//Get from Cache
	if t.todoCache.CacheLoad(listid) {
		result, err := t.todoCache.CacheGet(listid)
		if err != nil {
			return model.List{}, err
		}
		return result, nil
	}
	//Get from mysql
	result, err := t.todoRepo.GetOne(listid)
	if err != nil {
		return model.List{}, err
	}
	t.todoCache.CacheSet(listid, result.Title, result.Content)
	return result, nil
}
