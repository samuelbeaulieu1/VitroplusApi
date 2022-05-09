package controllers

import (
	"net/http"
	"time"

	"github.com/samuelbeaulieu1/gimlet"
	"github.com/samuelbeaulieu1/gimlet/responses"
	"github.com/samuelbeaulieu1/vitroplus-api/src/classes"
	"github.com/samuelbeaulieu1/vitroplus-api/src/dto"
	"github.com/samuelbeaulieu1/vitroplus-api/src/middlewares"
	"github.com/samuelbeaulieu1/vitroplus-api/src/services"
)

type ClockController struct{}

func NewClockController() *ClockController {
	return &ClockController{}
}

func (controller *ClockController) RegisterRoutes(router gimlet.Router) {
	router.Group("/Clock/", func(r gimlet.Router) {
		r.GET("Date/{date}/For/{employeeId}", controller.getEmployeeClocks).Use(middlewares.AuthenticateAdmin)
		r.GET("Between/{startDate}/{endDate}/In/{branchId}", controller.getBranchClocksInTimeframe).Use(middlewares.AuthenticateAdmin)
		r.GET("Between/{startDate}/{endDate}/For/{employeeId}", controller.getEmployeeClocksInTimeframe).Use(middlewares.AuthenticateAdmin)
		r.PUT("Date/{date}/For/{employeeId}", controller.updateEmployeeClocks).Use(middlewares.AuthenticateAdmin)
		r.POST("", controller.clockInOut)
		r.POST("Between/{startDate}/{endDate}/In/{branchId}/ToPDF", controller.getBranchReportPDF).Use(middlewares.AuthenticateAdmin)
	})
}

func getDateParam(key string, ctx *gimlet.Context) (*time.Time, bool) {
	dateStr := ctx.GetParam(key)
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		ctx.WriteJSONError(http.StatusBadRequest, responses.NewError("Date invalide"))
		return nil, false
	}

	return &date, true
}

func (controller *ClockController) getEmployeeClocks(ctx *gimlet.Context) {
	employeeId := ctx.GetParam("employeeId")
	date, ok := getDateParam("date", ctx)
	if !ok {
		return
	}

	if clocks, err := services.NewClockService().GetEmployeeClocks(employeeId, *date); err != nil {
		ctx.WriteJSONError(http.StatusBadRequest, err)
	} else {
		var clocksDto dto.EmployeeClocksDTO
		gimlet.ParseModelToDTO(&clocks, &clocksDto)
		ctx.WriteJSONResponse(clocksDto)
	}
}

func (controller *ClockController) getBranchClocksInTimeframe(ctx *gimlet.Context) {
	branchId := ctx.GetParam("branchId")
	startDate, startOk := getDateParam("startDate", ctx)
	endDate, endOk := getDateParam("endDate", ctx)
	if !startOk || !endOk {
		return
	}

	if report, err := services.NewClockService().GetBranchClocksInTimeframe(branchId, *startDate, *endDate); err != nil {
		ctx.WriteJSONError(http.StatusBadRequest, err)
	} else {
		var reportDto dto.BranchReportDTO
		gimlet.ParseModelToDTO(&report, &reportDto)
		ctx.WriteJSONResponse(reportDto)
	}
}

func (controller *ClockController) getBranchReportPDF(ctx *gimlet.Context) {
	branchId := ctx.GetParam("branchId")
	startDate, startOk := getDateParam("startDate", ctx)
	endDate, endOk := getDateParam("endDate", ctx)
	if !startOk || !endOk {
		return
	}
	var report classes.BranchReportRequest
	ctx.ParseBody(&report)
	report.StartDate = *startDate
	report.EndDate = *endDate

	branch, err := services.NewBranchService().Get(branchId)
	if err != nil {
		ctx.WriteJSONError(http.StatusBadRequest, err)
		return
	}

	report.Branch = branch
	if pdf, err := services.CreateReport(&report); err != nil {
		ctx.WriteJSONError(http.StatusBadRequest, err)
	} else {
		ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=rapport.pdf")
		ctx.Writer.Header().Set("Content-Type", "application/pdf")
		ctx.Writer.Write(pdf)
	}
}

func (controller *ClockController) getEmployeeClocksInTimeframe(ctx *gimlet.Context) {
	employeeId := ctx.GetParam("employeeId")
	startDate, startOk := getDateParam("startDate", ctx)
	endDate, endOk := getDateParam("endDate", ctx)
	if !startOk || !endOk {
		return
	}

	if clocks, err := services.NewClockService().GetEmployeeClocksInTimeframe(employeeId, *startDate, *endDate); err != nil {
		ctx.WriteJSONError(http.StatusBadRequest, err)
	} else {
		var clocksDto dto.ClocksReportDTO
		gimlet.ParseModelToDTO(&clocks, &clocksDto)
		ctx.WriteJSONResponse(clocksDto)
	}
}

func (controller *ClockController) updateEmployeeClocks(ctx *gimlet.Context) {
	var req classes.UpdateEmployeeClocksRequest
	ctx.ParseBody(&req)
	req.EmployeeID = ctx.GetParam("employeeId")
	date, ok := getDateParam("date", ctx)
	if !ok {
		return
	}
	req.Date = *date

	if err := services.NewClockService().UpdateEmployeeClocks(&req); err != nil {
		ctx.WriteJSONError(http.StatusBadRequest, err)
	} else {
		ctx.WriteJSONResponse(&responses.RequestResponseMessage{
			Message: "Les entrées de l'employé ont été modifiées",
		})
	}
}

func (controller *ClockController) clockInOut(ctx *gimlet.Context) {
	var req classes.ClockInRequest
	ctx.ParseBody(&req)

	if entry, err := services.NewClockService().ClockInOut(&req); err != nil {
		ctx.WriteJSONError(http.StatusBadRequest, err)
	} else {
		ctx.WriteJSONResponse(entry)
	}
}
