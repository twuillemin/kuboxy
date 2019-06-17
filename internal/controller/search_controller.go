package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/twuillemin/kuboxy/pkg/search"
)

func registerSearchControllers(e *echo.Echo) {

	e.POST("api/v1/search/:contextName", postSearch)
}

// postSearch searches the context for all kind objects
// @Summary Search objects
// @Description Search the context for all kind objects. All the parameters (except the object types) can be given as regexp.
// @ID post-search
// @Tags Search
// @Accept application/json
// @Produce application/json
// @Param contextName path string true "the name of the context"
// @Param body body search.Parameter true "the parameters of the search"
// @Success 200 {array} array interface{}
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/search/{contextName} [post]
func postSearch(e echo.Context) error {

	contextName := e.Param("contextName")

	// Parse the information from the body
	searchParameter := new(search.Parameter)
	if err := e.Bind(searchParameter); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Build the report
	results, err := search.Search(contextName, *searchParameter)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, results)
}
