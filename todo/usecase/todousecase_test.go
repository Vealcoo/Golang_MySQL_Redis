package usecase

import (
	"CleanTodo/model"
	"CleanTodo/model/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodoCache := new(mocks.TodoCache)
	t.Run("succes", func(t *testing.T) {
		mockTodo := model.List{
			Title:   "test",
			Content: "test",
		}
		mockTodoRepo.On("Create", mock.Anything, mock.Anything).Return(int64(1), nil).Once()

		u := NewTodoUsecase(mockTodoRepo, mockTodoCache)

		result, err := u.Create(mockTodo.Title, mockTodo.Content)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), result)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockTodo := model.List{
		Listid: "1",
	}
	t.Run("success", func(t *testing.T) {
		mockTodoRepo := new(mocks.TodoRepository)
		mockTodoCache := new(mocks.TodoCache)
		mockTodoCache.On("CacheLoad", mock.Anything).Return(false).Once()
		// mockTodoCache.On("CacheDel", mock.Anything).Return(nil).Once()
		mockTodoRepo.On("Delete", mock.Anything).Return(int64(1), nil).Once()

		u := NewTodoUsecase(mockTodoRepo, mockTodoCache)

		err := u.Delete(mockTodo.Listid)
		assert.NoError(t, err)
		mockTodoCache.AssertExpectations(t)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockTodo := model.List{
		Listid:  "1",
		Title:   "update",
		Content: "update",
	}
	t.Run("sucess", func(t *testing.T) {
		mockTodoRepo := new(mocks.TodoRepository)
		mockTodoCache := new(mocks.TodoCache)
		mockTodoCache.On("CacheLoad", mock.Anything).Return(false).Once()
		// mockTodoCache.On("CacheSet", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		mockTodoRepo.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(int64(1), nil).Once()

		u := NewTodoUsecase(mockTodoRepo, mockTodoCache)
		err := u.Update(mockTodo.Listid, mockTodo.Title, mockTodo.Content)
		assert.NoError(t, err)
		mockTodoCache.AssertExpectations(t)
		mockTodoRepo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	mockLists := []model.List{
		{
			Listid:  "1",
			Title:   "test",
			Content: "test",
		},
		{
			Listid:  "2",
			Title:   "test2",
			Content: "test2",
		},
	}
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodoCache := new(mocks.TodoCache)
	mockTodoRepo.On("GetAll").Return(mockLists, nil).Once()

	u := NewTodoUsecase(mockTodoRepo, mockTodoCache)
	result, err := u.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, mockLists, result)
	mockTodoRepo.AssertExpectations(t)
}

func TestGetOne(t *testing.T) {
	mockList := model.List{
		Listid:  "1",
		Title:   "test",
		Content: "test",
	}
	mockTodoRepo := new(mocks.TodoRepository)
	mockTodoCache := new(mocks.TodoCache)
	mockTodoCache.On("CacheLoad", mock.Anything).Return(false).Once()
	mockTodoCache.On("CacheSet", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	mockTodoRepo.On("GetOne", mock.Anything).Return(mockList, nil).Once()

	u := NewTodoUsecase(mockTodoRepo, mockTodoCache)
	result, err := u.GetOne(mockList.Listid)
	assert.NoError(t, err)
	assert.Equal(t, mockList, result)
	mockTodoRepo.AssertExpectations(t)
}
