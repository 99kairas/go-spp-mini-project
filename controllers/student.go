package controllers

import (
	"go-spp/middlewares"
	"go-spp/models/payloads"
	"go-spp/repositories"
	"go-spp/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterStudentController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	payloadUser := payloads.CreateStudentRequest{}
	c.Bind(&payloadUser)

	response, err := usecase.CreateStudent(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error create student",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success register student",
		Data:    response,
	})
}

func LoginStudentController(c echo.Context) error {
	payloadUser := payloads.LoginStudentRequest{}

	c.Bind(&payloadUser)

	response, err := usecase.LoginStudent(&payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error login student",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success login",
		Data:    response,
	})
}

func UpdatePasswordController(c echo.Context) error {
	payloadUser := payloads.UpdatePasswordRequest{}

	userId, err := middlewares.IsUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "this route only for user",
		})
	}

	c.Bind(&payloadUser)

	_, err = usecase.UpdatePassword(userId, &payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "failed change password",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success change password",
	})

}

func UpdateProfileController(c echo.Context) error {
	payloadUser := payloads.UpdateProfileStudentRequest{}

	c.Bind(&payloadUser)

	studentID, err := middlewares.IsUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "this route only for user",
		})
	}

	student, err := repositories.GetStudentByID(studentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: err.Error(),
		})
	}

	c.Bind(&student)

	response, err := usecase.UpdateProfile(student, &payloadUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success update profile",
		Data:    response,
	})
}
