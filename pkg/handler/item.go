package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhidAbdulazizov/todo-app/models"
	"net/http"
	"strconv"
)

// @Summary Create item in todo list
// @Security JWTTokenAuth
// @Tags items
// @Description create item in todo list
// @ID create-item-in-todo-list
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "list id"
// @Param input body models.TodoItem true "item info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/lists/:id/items/ [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponseError(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input models.TodoItem

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.TodoItem.Create(userId, listId, input)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All items is list
// @Security JWTTokenAuth
// @Tags items
// @Description get all items is list
// @ID get-all-items-is-list
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "list id"
// @Success 200 {object} []models.TodoItem
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/lists/:id/items/ [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponseError(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.service.TodoItem.GetAll(userId, listId)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get item by id
// @Security JWTTokenAuth
// @Tags items
// @Description get item by id
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "item id"
// @Success 200 {object} models.TodoItem
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/items/:id [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponseError(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.service.TodoItem.GetById(userId, itemId)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update item
// @Security JWTTokenAuth
// @Tags items
// @Description update list
// @ID update-item
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "item id"
// @Param input body models.UpdateItemInput true "update item info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/items/:id [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponseError(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input models.UpdateItemInput

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.TodoItem.Update(userId, id, input)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary Delete item
// @Security JWTTokenAuth
// @Tags items
// @Description delete item
// @ID delete-item
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "item id"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/items/:id [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponseError(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.service.TodoItem.Delete(userId, itemId)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
