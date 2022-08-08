package models

import (
	"database/sql"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Detail   string `json:"detail"`
	Assignee string `json:"assignee"`
	Due      string `json:"due"`
	Status   int    `json:"status"`
}

type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name, &task.Detail, &task.Assignee, &task.Due, &task.Status)

		if err2 != nil {
			panic(err2)
		}

		result.Tasks = append(result.Tasks, task)
	}
	return result
}

func PutTask(db *sql.DB, name string, detail string, assignee string, due string, status int) (int64, error) {

	sql := "INSERT INTO tasks(name, detail, assignee, due, status) VALUES (?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(name, detail, assignee, due, status)

	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

func EditTask(db *sql.DB, taskId int, name string, detail string, assignee string, due string, status int) (int64, error) {
	sql := "UPDATE tasks SET name = ?, detail =?, assignee = ?, due = ?, status = ? WHERE id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(name, detail, assignee, due, status, taskId)
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}
