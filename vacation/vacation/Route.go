package vacation

import (
	"github.com/labstack/echo"
	"net/http"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	vacation := &Vacation{}
	e.GET("/vacations", vacation.Get)
	e.GET("/vacations/:id", vacation.FindById)
	e.POST("/vacations", vacation.Persist)
	e.PUT("/vacations/:id", vacation.Put)
	e.DELETE("/vacations/:id", vacation.Remove)
	e.PUT("/vacations/cancel", Cancel)
	e.PUT("/vacations/approve", Approve)
	e.PUT("/vacations/confirmused", ConfirmUsed)
	e.PUT("/vacations/update", Update)
	vacationDaysLeft := &VacationDaysLeft{}
	e.GET("/vacationDaysLefts", vacationDaysLeft.Get)
	e.GET("/vacationDaysLefts/:id", vacationDaysLeft.FindById)
	e.POST("/vacationDaysLefts", vacationDaysLeft.Persist)
	e.PUT("/vacationDaysLefts/:id", vacationDaysLeft.Put)
	e.DELETE("/vacationDaysLefts/:id", vacationDaysLeft.Remove)
	e.PUT("/vacationDaysLefts/add", Add)
	e.PUT("/vacationDaysLefts/use", Use)
	vacationDaysStatus := &VacationDaysStatus{}
	e.GET("/vacationDaysStatuses", vacationDaysStatus.Get)
	e.GET("/vacationDaysStatuses/:id", vacationDaysStatus.FindById)
	e.POST("/vacationDaysStatuses", vacationDaysStatus.Persist)
	e.PUT("/vacationDaysStatuses/:id", vacationDaysStatus.Put)
	e.DELETE("/vacationDaysStatuses/:id", vacationDaysStatus.Remove)
	vacationStatus := &VacationStatus{}
	e.GET("/vacationStatuses", vacationStatus.Get)
	e.GET("/vacationStatuses/:id", vacationStatus.FindById)
	e.POST("/vacationStatuses", vacationStatus.Persist)
	e.PUT("/vacationStatuses/:id", vacationStatus.Put)
	e.DELETE("/vacationStatuses/:id", vacationStatus.Remove)
	vacationDaysLeft := &VacationDaysLeft{}
	e.GET("/vacationDaysLefts", vacationDaysLeft.Get)
	e.GET("/vacationDaysLefts/:id", vacationDaysLeft.FindById)
	e.POST("/vacationDaysLefts", vacationDaysLeft.Persist)
	e.PUT("/vacationDaysLefts/:id", vacationDaysLeft.Put)
	e.DELETE("/vacationDaysLefts/:id", vacationDaysLeft.Remove)
	return e
}
