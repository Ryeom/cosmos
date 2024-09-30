package router

import (
	"fmt"
	"github.com/Ryeom/cosmos/service/workspace"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CreateWorkspace(c echo.Context) error {
	result := GetDefaultResult()
	w := workspace.NewWorkspace()
	if bindErr := c.Bind(&w); bindErr != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	w.PrintDataInfo()
	err := w.Save()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, result.OK())
}

func UpdateWorkspace(c echo.Context) error {
	result := GetDefaultResult()
	w := workspace.NewWorkspace()
	if bindErr := c.Bind(&w); bindErr != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	//
	fmt.Println(w)
	return c.JSON(http.StatusOK, result.OK())
}

func DeleteWorkspace(c echo.Context) error {
	result := GetDefaultResult()
	w := workspace.NewWorkspace()
	if bindErr := c.Bind(&w); bindErr != nil {
		return c.JSON(http.StatusBadRequest, result)
	}
	//
	fmt.Println(w)
	return c.JSON(http.StatusOK, result.OK())
}

func GetWorkspace(c echo.Context) error {
	result := GetDefaultResult()
	w := workspace.NewWorkspace()
	if bindErr := c.Bind(&w); bindErr != nil {
		return c.JSON(http.StatusBadRequest, result)
	}

	fmt.Println(w)
	return c.JSON(http.StatusOK, result.OK())
}
