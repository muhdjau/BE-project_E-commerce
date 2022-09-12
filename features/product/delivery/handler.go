package delivery

import (
	"project/e-commerce/features/product"
	"project/e-commerce/middlewares"
	"project/e-commerce/utils/helper"

	"github.com/labstack/echo/v4"
)

type ProductDelivery struct {
	productUsecase product.UsecaseInterface
}

func New(e *echo.Echo, usecase product.UsecaseInterface) {
	handler := &ProductDelivery{
		productUsecase: usecase,
	}

	e.POST("/products", handler.PostData, middlewares.JWTMiddleware())

}

func (delivery *ProductDelivery) PostData(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	if idToken == 0 {
		return c.JSON(400, helper.FailedResponseHelper("Unauthorized"))
	}

	var dataProduct ProductRequest
	errBind := c.Bind(&dataProduct)
	if errBind != nil {
		return c.JSON(400, helper.FailedResponseHelper("error bind"))
	}

	row, err := delivery.productUsecase.PostData(toCore(dataProduct))
	if err != nil {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	if row != 1 {
		return c.JSON(500, helper.FailedResponseHelper("error insert data"))
	}
	return c.JSON(201, helper.SuccessResponseHelper("success insert data"))
}
