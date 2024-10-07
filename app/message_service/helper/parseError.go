package helper

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrNoUser            = errors.New("user doesn't exist")
	ErrNoData            = errors.New("no data found")
	ErrQuery             = errors.New("query execution failed")
	ErrInvalidId         = errors.New("invalid id")
	ErrInvalidDateFormat = errors.New("invalid date format")
	ErrBindJSON          = errors.New("unable to bind json")
	ErrUserExists        = errors.New("user already exist")

	ErrNoRows        = errors.New("no rows in result set")
	ErrScan          = errors.New("row scanning failed")
	ErrRowsAffected  = errors.New("unable to get affected row")
	ErrNoAffectedRow = errors.New("rows affected is 0")
	ErrLastInsertId  = errors.New("unable to get last insert id")
	ErrNoUpdate      = errors.New("data already exists")
	ErrParam         = errors.New("error or missing parameter")
	ErrCredential    = errors.New("password or email doesn't match")
)

func ParseError(err error, ctx echo.Context) error {
	status := http.StatusOK

	switch {
	case errors.Is(err, ErrQuery):
		fallthrough
	case errors.Is(err, ErrNoUser):
		status = http.StatusNotFound

	case errors.Is(err, ErrNoData):
		status = http.StatusNotFound

	case errors.Is(err, ErrParam):
		status = http.StatusBadRequest

	case errors.Is(err, ErrBindJSON):
		status = http.StatusBadRequest

	case errors.Is(err, ErrInvalidId):
		status = http.StatusBadRequest

	case errors.Is(err, ErrCredential):
		status = http.StatusBadRequest

	case errors.Is(err, ErrInvalidDateFormat):
		status = http.StatusBadRequest

	case errors.Is(err, ErrUserExists):
		status = http.StatusBadRequest

	case errors.Is(err, ErrNoUpdate):
		status = http.StatusBadRequest

	default:
		status = http.StatusInternalServerError

	}

	return ctx.JSON(status, map[string]interface{}{"message": err.Error()})
}
