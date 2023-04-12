package controller

import (
	"net/http"
	"rest-api-employee/config"
	"rest-api-employee/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AddEmployee(c echo.Context) error {
	var employee model.Employee
	c.Bind(&employee)

	result := config.DB.Create(&employee)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "Gagal menambahkan Employeee", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "Berhasil menambahkan Employee", Data: employee,
	})
}

func GetEmployees(c echo.Context) error {
	var employees []model.Employee

	result := config.DB.Find(&employees)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "Employees tidak ada", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "success", Data: employees,
	})
}

func GetEmployeeById(c echo.Context) error {
	id := c.Param("id")
	var employee model.Employee
	if err := config.DB.First(&employee, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, model.Response{
				Message: "Employee tidak ada", Data: nil,
			})
		}
	}
	return c.JSON(http.StatusOK, model.Response{
		Message: "success", Data: employee,
	})
}

func UpdateEmployee(c echo.Context) error {
	id := c.Param("id")
	var employee model.Employee
	if err := config.DB.First(&employee, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, model.Response{
				Message: "Employee tidak ada", Data: nil,
			})
		}
	}
	var updatedEmployee model.Employee
	if err := c.Bind(&updatedEmployee); err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "request body salah", Data: nil,
		})
	}
	employee.Firstname = updatedEmployee.Firstname
	employee.Lastname = updatedEmployee.Lastname
	employee.Email = updatedEmployee.Email
	employee.Salary = updatedEmployee.Salary

	if err := config.DB.Save(&employee).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "gagal mengubah employee", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "success", Data: employee,
	})
}

func DeleteEmployee(c echo.Context) error {
	id := c.Param("id")

	var employee model.Employee
	if err := config.DB.First(&employee, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "Employee tidak ada", Data: nil,
		})
	}

	if err := config.DB.Delete(&employee).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Message: "gagal menghapus employee", Data: nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Message: "success", Data: employee,
	})
}
