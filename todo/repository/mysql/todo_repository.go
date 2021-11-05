package mysql

import (
	"CleanTodo/model"
	"database/sql"
)

type mysqlTodoRepository struct {
	db *sql.DB
}

func NewMysqlTodoRepository(c *sql.DB) model.TodoRepository {
	return &mysqlTodoRepository{
		db: c,
	}
}

func (m *mysqlTodoRepository) Create(title, content string) (int64, error) {
	result, err := m.db.Exec("INSERT INTO list (title, content) values (?, ?)", title, content)
	if err != nil {
		return 0, err
	}
	listid, err := result.LastInsertId()
	return listid, nil
}

func (m *mysqlTodoRepository) Delete(listid string) (int64, error) {
	result, err := m.db.Exec("DELETE FROM list WHERE listid=?", listid)
	if err != nil {
		return 0, err
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsaffected, nil
}

func (m *mysqlTodoRepository) Update(listid, title, content string) (int64, error) {
	result, err := m.db.Exec("UPDATE list SET title=?, content=? WHERE listid=?", title, content, listid)
	if err != nil {
		return 0, err
	}
	rowsaffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsaffected, nil
}

func (m *mysqlTodoRepository) GetAll() ([]model.List, error) {
	rows, err := m.db.Query("SELECT listid, title, content FROM list")
	defer rows.Close()
	if err != nil {

	}
	var lists []model.List
	for rows.Next() {
		var list model.List
		err = rows.Scan(&list.Listid, &list.Title, &list.Content)
		if err != nil {

		}
		lists = append(lists, list)
	}
	return lists, nil
}

func (m *mysqlTodoRepository) GetOne(listid string) (model.List, error) {
	var list model.List
	row := m.db.QueryRow("SELECT listid, title, content FROM list WHERE listid=" + listid)
	if err := row.Scan(&list.Listid, &list.Title, &list.Content); err != nil {
		return list, err
	}
	return list, nil
}
