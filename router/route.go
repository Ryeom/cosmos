package router

import (
	"github.com/Ryeom/cosmos/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Initialize(e *echo.Echo) {
	serviceApi := e.Group("/cosmos")
	{
		route(serviceApi)
	}
}
func route(g *echo.Group) {
	g.GET("/healthcheck", healthCheck)

	g.POST("/workspace", CreateWorkspace)
	g.PUT("/workspace", UpdateWorkspace)
	g.GET("/workspace", GetWorkspace)
	g.DELETE("/workspace", DeleteWorkspace)

	g.POST("/schedule", CreateSchedule)
	g.PUT("/schedule", UpdateSchedule)
	g.GET("/schedule", GetSchedule)
	g.DELETE("/schedule", DeleteSchedule)

}

func healthCheck(c echo.Context) error {
	result := HttpResult{
		ResultCode: 200,
		ResultMsg:  "OK",
		ResultData: map[string]interface{}{
			"uptime":  log.UpTime,
			"version": 0,
		},
	}

	return c.JSON(http.StatusOK, result)
}

func Cors(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		corsConfig := middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowHeaders:     []string{"Accept", "Cache-Control", "Content-Type", "X-Requested-With"},
			AllowCredentials: true,
		}
		CORSMiddleware := middleware.CORSWithConfig(corsConfig)
		CORSHandler := CORSMiddleware(next)
		return CORSHandler(ctx)
	}
}
