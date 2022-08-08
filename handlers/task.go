package handlers

import (
	"database/sql"
	"net/http"
	"simpletask-backend/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task
		c.Bind(&task)

		id, err := models.PutTask(db, task.Name, task.Detail, task.Assignee, task.Due, task.Status)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, H{
			"created": id,
		})
	}
}

func EditTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var task models.Task
		task.ID = id
		c.Bind(&task)

		_, err := models.EditTask(db, task.ID, task.Name, task.Detail, task.Assignee, task.Due, task.Status)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, H{
			"updated": task,
		})

	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := models.DeleteTask(db, id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, H{
			"deleted": id,
		})

	}
}
