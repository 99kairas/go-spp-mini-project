package controllers

import (
	"go-spp/middlewares"
	"go-spp/models/payloads"
	"go-spp/usecase"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func AdminCreatePaymentByStudentIDController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	payloadPayment := payloads.AdminCreatePaymentByStudentIDRequest{}
	c.Bind(&payloadPayment)

	response, err := usecase.CreatePaymentByStudentID(&payloadPayment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error create payment",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success create payment",
		Data:    response,
	})
}

func AdminCreatePaymentAllStudentController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	payloadPayment := payloads.AdminCreatePaymentAllStudentRequest{}
	c.Bind(&payloadPayment)

	response, err := usecase.CreatePaymentAllStudent(&payloadPayment)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "error create payment",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success create payment",
		Data:    response,
	})
}

func GetAllPaymentsController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	payment, err := usecase.GetAllPayments()
	if err != nil {
		return c.JSON(http.StatusNotFound, payloads.Response{
			Message: "payment not found",
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success get all data",
		Data:    payment,
	})
}

func GetPaymentByIDController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	paymentID := c.Param("id")

	paymentUUID, err := uuid.Parse(paymentID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "invalid payment id format",
		})
	}

	payment, err := usecase.GetPaymentByID(paymentUUID)
	if err != nil {
		return c.JSON(http.StatusNotFound, payloads.Response{
			Message: "success not found",
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success get payment",
		Data:    payment,
	})
}

func GetPaymentsWithPhotoController(c echo.Context) error {
	if _, err := middlewares.IsAdmin(c); err != nil {
		return c.JSON(http.StatusUnauthorized, payloads.Response{
			Message: "route only for admin",
		})
	}

	payment, err := usecase.GetPaymentsWithPhoto()
	if err != nil {
		return c.JSON(http.StatusNotFound, payloads.Response{
			Message: "payment not found",
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "success get all data",
		Data:    payment,
	})
}

func ApprovePaymentController(c echo.Context) error {
	var request payloads.ApproveRejectPaymentRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "invalid request body",
			Data:    err.Error(),
		})
	}

	response, err := usecase.ApprovePayment(request.PaymentID, request.AdminID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, payloads.Response{
			Message: "failed to approve payment",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, payloads.Response{
		Message: "payment approved successfully",
		Data:    response,
	})
}
