package controllers

import (
	"fmt"
	"go-spp/models/payloads"
	"go-spp/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChatBot struct {
	Question string `json:"question"`
}

func OpenAIController(c echo.Context) error {
	req := new(ChatBot)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "invalid request")
	}

	recommendation, err := utils.GetChatBot(req.Question)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to get chat from bot")
	}

	response := fmt.Sprintf("here's the answer from bot : %s", recommendation)

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success get response",
		Data:    response,
	})

}
