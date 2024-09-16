package handlers

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GetUser godoc
// @Summary Get a user
// @Description Get a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id query int true "User ID"
// @Param name query string false "User Name"
// @Security ApiKeyAuth
// @Success 200 {object} User
// @Failure 401 {object} string "Unauthorized"
// @Router /user [get]
func GetUser(ctx *fasthttp.RequestCtx) {
	// 检查认证
	token := string(ctx.Request.Header.Peek("T"))
	if token == "" {
		ctx.SetStatusCode(fasthttp.StatusUnauthorized)
		ctx.SetBodyString("Unauthorized: Missing token")
		return
	}
	// 这里可以添加更多的token验证逻辑

	id := ctx.QueryArgs().GetUintOrZero("id")
	name := string(ctx.QueryArgs().Peek("name"))
	user := User{
		ID:   id,
		Name: name,
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(jsonUser)
}
