package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vakhidAbdulazizov/todo-app/models"
	"net/http"
)

type signInInput struct {
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type forgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

type restorePasswordInput struct {
	Email      string `json:"email" binding:"required"`
	ConfirmKey string `json:"confirmKey" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description create account
// @ID sign-in-account
// @Accept  json
// @Produce  json
// @Param input body signInInput true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /auth/sign-in [post]
func (h *Handler) singIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input.UserName, input.Password)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /auth/sign-up [post]
func (h *Handler) singUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Authorization.CreateUser(input)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) forgotPass(c *gin.Context) {
	var input forgotPasswordInput

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Authorization.ForgotPassword(input.Email)

	if err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) restorePassword(c *gin.Context) {
	var input restorePasswordInput

	if err := c.BindJSON(&input); err != nil {
		newResponseError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.RestorePassword(input.Email, input.ConfirmKey, input.Password)

	if err != nil {
		newResponseError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
