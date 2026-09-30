package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/n1tro-dev/todo-my"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
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

type getAllListResponse struct {
	Data []todo.TodoList
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		return
	}

	lists, err := h.service.TodoList.GetAll(userId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: lists,
	})

}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := h.GetUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.service.TodoList.GetById(userId, listId)
	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
