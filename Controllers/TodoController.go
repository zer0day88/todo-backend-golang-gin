package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo/backend/Models"
	"todo/backend/Repository"
)

func GetTodos(c *gin.Context) {
	var todos []Models.Todo
	var todo Models.Todo
	var resp Models.Response
	c.Bind(&todo)

	if todo.Activity_group_id == 0 {
		err := Repository.GetAllTodos(&todos)
		if err != nil {
			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = make([]string, 0)
			c.JSON(http.StatusNotFound, resp)
		} else {
			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = todos
			c.JSON(http.StatusOK, resp)
		}
	}

	if todo.Activity_group_id > 0 {

		id := strconv.Itoa(todo.Activity_group_id)
		err := Repository.GetAllTodosByActivityId(&todos, id)
		if err != nil {
			resp.Status = "Success"
			resp.Message = "Success"
			resp.Data = make([]string, 0)
			c.JSON(http.StatusOK, resp)
		} else {

			if todos == nil {
				resp.Status = "Success"
				resp.Message = "Success"
				resp.Data = make([]string, 0)
				c.JSON(http.StatusNotFound, resp)
			} else {
				resp.Status = "Success"
				resp.Message = "Success"
				resp.Data = todos
				c.JSON(http.StatusOK, resp)
			}
		}
	}
}

func CreateTodo(c *gin.Context) {
	var todo Models.Todo
	var resp Models.Response
	var todo_created Models.Todo_Created
	c.BindJSON(&todo)
	if todo.Title == "" {
		resp.Status = "Bad Request"
		resp.Message = "title cannot be null"
		resp.Data = new(Models.EmptyObject)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if todo.Activity_group_id == 0 {
		resp.Status = "Bad Request"
		resp.Message = "activity_group_id cannot be null"
		resp.Data = new(Models.EmptyObject)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err := Repository.CreateTodo(&todo)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		todo_created.Id = todo.Id
		todo_created.Activity_group_id = todo.Activity_group_id
		todo_created.Title = todo.Title
		todo_created.Priority = todo.Priority
		todo_created.Is_active = todo.Is_active
		todo_created.Created_at = todo.Created_at
		todo_created.Updated_at = todo.Updated_at

		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = todo_created
		c.JSON(http.StatusCreated, resp)
	}
}

func GetTodoById(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Models.Todo
	var resp Models.Response
	err := Repository.GetTodoById(&todo, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Todo with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.JSON(http.StatusNotFound, resp)
	} else {

		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = todo
		c.JSON(http.StatusOK, resp)
	}
}

func UpdateTodo(c *gin.Context) {
	var todo Models.Todo
	var todo_view Models.Todo
	var resp Models.Response

	id := c.Params.ByName("id")

	c.BindJSON(&todo)

	err := Repository.GetTodoById(&todo_view, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Todo with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.JSON(http.StatusNotFound, resp)
		return
	}

	err = Repository.UpdateTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		_ = Repository.GetTodoById(&todo_view, id)

		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = todo_view
		c.JSON(http.StatusOK, resp)
	}
}

func DeleteTodo(c *gin.Context) {
	var todo Models.Todo
	var resp Models.Response
	id := c.Params.ByName("id")
	err := Repository.GetTodoById(&todo, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Todo with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.JSON(http.StatusNotFound, resp)
		return
	}

	err = Repository.DeleteTodo(&todo, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Todo with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.JSON(http.StatusNotFound, resp)
	} else {

		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = new(Models.EmptyObject)
		c.JSON(http.StatusOK, resp)
	}
}
