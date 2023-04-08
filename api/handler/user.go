package handler

import (
	"app/api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @ID register_user
// @Router /register [POST]
// @Summary Register User
// @Description Register User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "CreateUserRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) Register(c *gin.Context) {

	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create brand", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Auth().Register(context.Background(), &createUser)
	if err != nil {
		h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create brand", http.StatusCreated, id)
}

// Login godoc
// @ID Login
// @Router /login [POST]
// @Summary Register User
// @Description Login User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.Login1 true "LoginRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) Login(c *gin.Context) {

	var loginData models.Login1

	err := c.ShouldBindJSON(&loginData) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create brand", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Auth().Login(context.Background(), &loginData)
	if err != nil {
		h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create brand", http.StatusCreated, id)
}
