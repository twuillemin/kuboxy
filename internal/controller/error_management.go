package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/twuillemin/kuboxy/pkg/context"
	"net/http"
)

func getHTTPError(err error) *echo.HTTPError {
	if e, ok := err.(*context.NotFoundError); ok {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("the context \"%s\" does not exist", e.ContextName()))
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
