package controllers

import (
	"go-spp/middlewares"
	"go-spp/models/payloads"
	"go-spp/repositories"
	"go-spp/usecase"
	"go-spp/utils"
	"net/http"

	"github.com/google/uuid"
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

	studentID, err := middlewares.IsUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "this route only for user",
		})
	}

	c.Bind(&payloadUser)

	_, err = usecase.UpdatePassword(studentID, &payloadUser)
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

func UploadPaymentPhotoController(c echo.Context) error {
	payloadUser := payloads.UploadPaymentPhotoRequest{}

	c.Bind(&payloadUser)

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error upload image form file",
		})
	}

	resp, err := utils.UploadPaymentPhoto(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error upload image to cloudinary",
		})
	}

	_, err = usecase.UploadImage(payloadUser.PaymentID, resp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error upload image to database",
			Data:    resp,
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "upload image success",
		Data:    resp,
	})
}

func GetPaymentsController(c echo.Context) error {
	studentID, err := middlewares.IsUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "this route is only for user",
		})
	}

	user, err := usecase.GetPayments(studentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success get all data",
		Data:    user,
	})
}

func GetDetailPaymentController(c echo.Context) error {
	studentID, err := middlewares.IsUser(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for authorized users",
		})
	}

	paymentID := c.Param("id")

	paymentUUID, err := uuid.Parse(paymentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "invalid payment id format",
		})
	}

	payment, err := usecase.GetDetailPayments(paymentUUID)
	if err != nil {
		return c.JSON(http.StatusNotFound, payloads.Response{
			Message: "payment not found",
		})
	}

	if studentID != payment[0].Student.ID {
		return c.JSON(http.StatusForbidden, payloads.Response{
			Message: "access denied",
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success get payment",
		Data:    payment,
	})
}

func DeleteStudentController(c echo.Context) error {
	studentID := c.Param("id")

	studentUUID, err := uuid.Parse(studentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "invalid student ID format",
		})
	}

	err = usecase.DeleteStudent(studentUUID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "failed to delete student",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "student deleted successfully",
	})
}
