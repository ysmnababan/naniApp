package grpchandler

import (
	"user_service/usecase"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecaseI
}

func (h *UserHandler) Login(e echo.Context) error {
	return nil
}
