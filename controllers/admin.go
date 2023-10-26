package controllers

import (
	"fmt"
	"go-spp/middlewares"
	"go-spp/models/payloads"
	"go-spp/usecase"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func RegisterAdminController(c echo.Context) error {
	payloadUser := payloads.CreateAdminRequest{}
	c.Bind(&payloadUser)

	response, err := usecase.CreateAdmin(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error create admin",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success register admin",
		Data:    response,
	})
}

func LoginAdminController(c echo.Context) error {
	payloadUser := payloads.LoginAdminRequest{}

	c.Bind(&payloadUser)

	response, err := usecase.LoginAdmin(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error login admin",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success login as admin",
		Data:    response,
	})
}

func GetStudentIDController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	studentID := c.Param("studentID")

	studentUUID, err := uuid.Parse(studentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "Invalid student ID format",
		})
	}

	student, err := usecase.GetStudentByID(studentUUID)
	if err != nil {
		return c.JSON(http.StatusNotFound, payloads.Response{
			Message: "Student not found",
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: fmt.Sprintf("Success get profile %s %s", student.FirstName, student.LastName),
		Data:    student,
	})
}
