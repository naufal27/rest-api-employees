package route

import (
	"rest-api-employee/controller"

	"github.com/labstack/echo/v4"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	e.POST("/employee", controller.AddEmployee)
	e.GET("/employee", controller.GetEmployees)
	e.GET("/employee/:id", controller.GetEmployeeById)
	e.PUT("/employee/:id", controller.UpdateEmployee)
	e.DELETE("/employee/:id", controller.DeleteEmployee)
	return e
}
