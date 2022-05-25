package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/backend/Models"
	"todo/backend/Repository"
)

func GetActivities(c *gin.Context) {

	var activity []Models.Activity
	var activity_view []Models.Activity_model
	var resp Models.Response
	err := Repository.GetAllActivities(&activity)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		for _, s := range activity {
			var act Models.Activity_model
			act.Id = s.Id
			act.Email = s.Email
			act.Title = s.Title
			act.Created_at = s.Created_at.Format("2006-01-02T03:04:05.000Z")
			act.Updated_at = s.Updated_at.Format("2006-01-02T03:04:05.000Z")
			if s.Deleted_at == nil {
				act.Deleted_at = nil
			} else {
				var deleted_time = s.Deleted_at.Format("2006-01-02T03:04:05.000Z")
				act.Deleted_at = &deleted_time
			}
			activity_view = append(activity_view, act)
		}

		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = activity_view
		c.JSON(http.StatusOK, resp)
	}
}

func CreateActivity(c *gin.Context) {
	var activity Models.Activity
	var resp Models.Response
	var activity_view Models.Activity_Created
	c.BindJSON(&activity)

	if activity.Title == "" {
		resp.Status = "Bad Request"
		resp.Message = "title cannot be null"
		resp.Data = new(Models.EmptyObject)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err := Repository.CreateActivity(&activity)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		activity_view.Id = activity.Id
		activity_view.Email = activity.Email
		activity_view.Title = activity.Title
		activity_view.Created_at = activity.Created_at
		activity_view.Updated_at = activity.Updated_at

		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = activity_view

		c.JSON(http.StatusCreated, resp)
	}
}

func GetActivityById(c *gin.Context) {
	id := c.Params.ByName("id")
	var activity Models.Activity
	var resp Models.Response
	err := Repository.GetActivityById(&activity, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Activity with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.JSON(http.StatusNotFound, resp)
	} else {

		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = activity
		c.JSON(http.StatusOK, resp)
	}
}

func UpdateActivity(c *gin.Context) {
	var activity Models.Activity
	var act Models.Activity
	var resp Models.Response

	id := c.Params.ByName("id")

	c.BindJSON(&activity)
	if activity.Title == "" {
		resp.Status = "Bad Request"
		resp.Message = "title cannot be null"
		resp.Data = new(Models.EmptyObject)

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err := Repository.GetActivityById(&act, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Activity with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.JSON(http.StatusNotFound, resp)
		return
	}

	err = Repository.UpdateActivity(&activity, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		_ = Repository.GetActivityById(&act, id)

		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = act
		c.JSON(http.StatusOK, resp)
	}
}

func DeleteActivity(c *gin.Context) {
	var activity Models.Activity
	var resp Models.Response
	id := c.Params.ByName("id")

	err := Repository.GetActivityById(&activity, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Activity with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.JSON(http.StatusNotFound, resp)
		return
	}

	err = Repository.DeleteActivity(&activity, id)
	if err != nil {
		resp.Status = "Not Found"
		resp.Message = fmt.Sprintf("Activity with ID %s Not Found", id)
		resp.Data = new(Models.EmptyObject)
		c.JSON(http.StatusNotFound, resp)
	} else {
		resp.Status = "Success"
		resp.Message = "Success"
		resp.Data = new(Models.EmptyObject)
		c.JSON(http.StatusOK, resp)
	}
}
