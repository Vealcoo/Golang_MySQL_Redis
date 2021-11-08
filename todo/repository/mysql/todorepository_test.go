package mysql

import (
	"CleanTodo/model"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCreate(t *testing.T) {
	input := &model.List{
		Title:   "懇親會",
		Content: "參加兒子懇親會",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	query := "^INSERT INTO list*"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(input.Title, input.Content).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewMysqlTodoRepository(db)

	result, err := a.Create(input.Title, input.Content)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result)
}

func TestDelete(t *testing.T) {
	input := &model.List{
		Listid: "1",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	query := "DELETE FROM list WHERE listid=?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(input.Listid).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewMysqlTodoRepository(db)

	result, err := a.Delete(input.Listid)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result)
}

func TestUpdate(t *testing.T) {
	input := &model.List{
		Listid:  "1",
		Title:   "update",
		Content: "update",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	query := "UPDATE list SET title=\\?, content=\\? WHERE listid=\\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(input.Title, input.Content, input.Listid).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewMysqlTodoRepository(db)

	result, err := a.Update(input.Listid, input.Title, input.Content)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), result)
}

func TestGetAll(t *testing.T) {
	mockLists := []model.List{
		model.List{
			Listid:  "1",
			Title:   "test",
			Content: "test",
		},
		model.List{
			Listid:  "2",
			Title:   "test2",
			Content: "test2",
		},
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"listid", "title", "conent"}).
		AddRow(mockLists[0].Listid, mockLists[0].Title, mockLists[0].Content).
		AddRow(mockLists[1].Listid, mockLists[1].Title, mockLists[1].Content)

	query := "SELECT listid, title, content FROM list"
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := NewMysqlTodoRepository(db)

	result, err := a.GetAll()
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestGetOne(t *testing.T) {
	mockList := model.List{
		Listid:  "1",
		Title:   "test",
		Content: "test",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"listid", "title", "conent"}).
		AddRow(mockList.Listid, mockList.Title, mockList.Content)

	query := "SELECT listid, title, content FROM list WHERE listid="
	mock.ExpectQuery(query).WillReturnRows(rows)

	a := NewMysqlTodoRepository(db)

	result, err := a.GetOne(mockList.Listid)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
