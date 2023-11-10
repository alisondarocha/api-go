package models

import (
	"api-go/context"
)

func Get(id int64) (todo Todo, err error) {
	appContext := context.GetAppContext()
	conn := appContext.Connection

	row := conn.QueryRow(`SELECT * FROM todos WHERE id=$1`, id)

	err = row.Scan(todo.Id, todo.Title, todo.Description, todo.Done)

	return
}
