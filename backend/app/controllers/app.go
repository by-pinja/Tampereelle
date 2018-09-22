package controllers

import (
	"github.com/revel/revel"
)

type JSONResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Test() revel.Result {
	response := JSONResponse{Success: true}
	return c.RenderJSON(response)
}
