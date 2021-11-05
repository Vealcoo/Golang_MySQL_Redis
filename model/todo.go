package model

type List struct {
	Listid  string
	Title   string
	Content string
}

type TodoUsecase interface {
	Create(title, content string) (int64, error)
	Delete(listid string) error
	Update(listid, title, content string) error
	GetAll() ([]List, error)
	GetOne(listid string) (List, error)
}

type TodoRepository interface {
	Create(title, content string) (int64, error)
	Delete(listid string) (int64, error)
	Update(listid, title, content string) (int64, error)
	GetAll() ([]List, error)
	GetOne(listid string) (List, error)
}

type TodoCache interface {
	CacheSet(listid, title, content string) error
	CacheDel(listid string) error
	CacheGet(listid string) (List, error)
	CacheLoad(listid string) bool
}
