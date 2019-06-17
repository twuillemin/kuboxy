package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/twuillemin/kuboxy/pkg/report"
)

func registerSummaryControllers(e *echo.Echo) {

	e.GET("api/v1/summary/:contextName", getSummary)
}

// getSummary generates a JSON representation of all the information in the given configuration
// @Summary Get the global status, or summary, of the given configuration
// @Description get the summary of a configuration
// @ID get-summary
// @Tags Summary
// @Produce application/json
// @Param contextName path string true "the name of the context"
// @Success 200 {object} report.ClusterStateReport
// @Failure 404 {object} HTTPError
// @Failure 500 {object} HTTPError
// @Router /api/v1/summary/{contextName} [get]
func getSummary(e echo.Context) error {

	contextName := e.Param("contextName")

	// Build the report
	stateReport, err := report.BuildReport(contextName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return e.JSON(http.StatusOK, stateReport)
}
