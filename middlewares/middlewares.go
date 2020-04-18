package middlewares

import (
	"Demo-RestApi/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

//middleware customcontext
func MidCustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &model.CustomContext{c}
		return next(cc)
	}
}

func ValidateUser(username, password string, c echo.Context) (bool, error) {
	if username == "sakib" && password == "mulla" {
		return true, nil
	}
	return false, nil
}

func CorsPolicy()  middleware.CORSConfig {
	mid := middleware.CORSConfig{
		AllowOrigins: []string{"10.0.2.102"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}
	return mid
}

