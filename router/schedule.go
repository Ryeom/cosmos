package router

import (
	"fmt"
	"github.com/Ryeom/cosmos/service/workspace"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateSchedule(c echo.Context) error {
	result := GetDefaultResult()
	w := workspace.NewWorkspace()
	if bindErr := c.Bind(&w); bindErr != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	//
	fmt.Println(w)
	return c.JSON(http.StatusOK, result.OK())
}

func UpdateSchedule(c echo.Context) error {
	result := GetDefaultResult()
	w := workspace.NewWorkspace()
	if bindErr := c.Bind(&w); bindErr != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	//
	fmt.Println(w)
	return c.JSON(http.StatusOK, result.OK())
}

func GetSchedule(c echo.Context) error {
	result := GetDefaultResult()
	w := workspace.NewWorkspace()
	if bindErr := c.Bind(&w); bindErr != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	//
	fmt.Println(w)
	return c.JSON(http.StatusOK, result.OK())
}
func DeleteSchedule(c echo.Context) error {
	result := GetDefaultResult()
	w := workspace.NewWorkspace()
	if bindErr := c.Bind(&w); bindErr != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	//
	fmt.Println(w)
	return c.JSON(http.StatusOK, result.OK())
}
