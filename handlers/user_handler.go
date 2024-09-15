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
// @Success 200 {object} User
// @Router /user [get]
func GetUser(ctx *fasthttp.RequestCtx) {
	id := ctx.QueryArgs().GetUintOrZero("id")
	name := ctx.QueryArgs().Peek("name")
	user := User{
		ID:   id,
		Name: string(name),
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
