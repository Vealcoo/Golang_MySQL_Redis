package mocks

import (
	"CleanTodo/model"

	"github.com/stretchr/testify/mock"
)

//TodoRepository Mock----------------------------------------------------------
type TodoRepository struct {
	mock.Mock
}

func (t *TodoRepository) Create(title, content string) (int64, error) {
	ret := t.Called(title, content)

	var r error
	if result, ok := ret.Get(1).(func(string, string) (int64, error)); ok {
		_, r = result(title, content)
	} else {
		_, r = 0, ret.Error(1)
	}
	return int64(1), r
}

func (t *TodoRepository) Delete(listid string) (int64, error) {
	ret := t.Called(listid)

	var r error
	if result, ok := ret.Get(1).(func(string) (int64, error)); ok {
		_, r = result(listid)
	} else {
		_, r = 0, ret.Error(1)
	}
	return int64(1), r
}

func (t *TodoRepository) Update(listid, title, content string) (int64, error) {
	ret := t.Called(listid, title, content)

	var r error
	if result, ok := ret.Get(1).(func(string, string, string) (int64, error)); ok {
		_, r = result(listid, title, content)
	} else {
		_, r = 0, ret.Error(1)
	}
	return int64(1), r
}

func (t *TodoRepository) GetAll() ([]model.List, error) {
	ret := t.Called()

	var l []model.List
	var r error
	if result, ok := ret.Get(0).(func() ([]model.List, error)); ok {
		l, r = result()
	} else {
		l, r = ret.Get(0).([]model.List), ret.Error(1)
	}
	return l, r
}

func (t *TodoRepository) GetOne(listid string) (model.List, error) {
	ret := t.Called(listid)

	var l model.List
	var r error
	if result, ok := ret.Get(0).(func(string) (model.List, error)); ok {
		l, r = result(listid)
	} else {
		l, r = ret.Get(0).(model.List), ret.Error(1)
	}
	return l, r
}

//TodoCache Mock-----------------------------------------------------------------
type TodoCache struct {
	mock.Mock
}

func (t *TodoCache) CacheSet(listid, title, content string) error {
	ret := t.Called(listid, title, content)

	var r error
	if result, ok := ret.Get(0).(func(string, string, string) error); ok {
		r = result(listid, title, content)
	} else {
		r = ret.Error(0)
	}
	return r
}

func (t *TodoCache) CacheGet(listid string) (model.List, error) {
	return model.List{}, nil
}

func (t *TodoCache) CacheDel(listid string) error {
	ret := t.Called(listid)

	var r error
	if result, ok := ret.Get(0).(func(string) error); ok {
		r = result(listid)
	} else {
		r = ret.Error(0)
	}
	return r
}

func (t *TodoCache) CacheLoad(listid string) bool {
	ret := t.Called(listid)

	var r bool
	if result, ok := ret.Get(0).(func(string) bool); ok {
		r = result(listid)
	} else {
		r = false
	}
	return r
}
