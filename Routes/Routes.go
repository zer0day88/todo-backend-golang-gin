package Routes

import (
	"github.com/gin-gonic/gin"
	"todo/backend/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	act := r.Group("/activity-groups")
	{
		act.GET("", Controllers.GetActivities)
		act.POST("", Controllers.CreateActivity)
		act.GET(":id", Controllers.GetActivityById)
		act.PATCH(":id", Controllers.UpdateActivity)
		act.DELETE(":id", Controllers.DeleteActivity)
	}

	todo := r.Group("/todo-items")
	{
		todo.GET("", Controllers.GetTodos)
		todo.POST("", Controllers.CreateTodo)
		todo.GET(":id", Controllers.GetTodoById)
		todo.PATCH(":id", Controllers.UpdateTodo)
		todo.DELETE(":id", Controllers.DeleteTodo)
	}
	return r
}
