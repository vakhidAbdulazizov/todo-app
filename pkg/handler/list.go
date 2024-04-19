package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhidAbdulazizov/todo-app/models"
	"net/http"
	"strconv"
)

type GetAllListResponse struct {
	Data []models.TodoList `json:"data"`
}

// @Summary Create todo list
// @Security JWTTokenAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body models.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	var input models.TodoList

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.service.TodoList.Create(userId, input)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Lists
// @Security JWTTokenAuth
// @Tags lists
// @Description get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} GetAllListResponse
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	lists, err := h.service.TodoList.GetAll(userId)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAllListResponse{
		Data: lists,
	})
}

// @Summary Get List  by ID
// @Security JWTTokenAuth
// @Tags lists
// @Description get list  by id
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "list id"
// @Success 200 {object} models.TodoList
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/lists/:id [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponseError(c, http.StatusBadRequest, "invalid id")
		return
	}

	item, err := h.service.TodoList.GetByID(userId, id)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)

}

// @Summary Update list
// @Security JWTTokenAuth
// @Tags lists
// @Description update list
// @ID update-list
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "list id"
// @Param input body models.UpdateListInput true "update list info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/lists/:id [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponseError(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input models.UpdateListInput

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.TodoList.Update(userId, id, input)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete list
// @Security JWTTokenAuth
// @Tags lists
// @Description delete list
// @ID delete-list
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "list id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/lists/:id [delete]
func (h *Handler) deleteList(c *gin.Context) {

	userId, err := getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponseError(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.service.TodoList.Delete(userId, id)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
