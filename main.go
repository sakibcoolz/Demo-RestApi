package main

import (
	"Demo-RestApi/handlers"
	"Demo-RestApi/middlewares"
	"Demo-RestApi/db"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
var MysqlPool *sql.DB

func init() {
	fmt.Println("Init")
	MysqlPool = db.DbConn()
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// Echo instance
	e := echo.New()

	//Use BasicAuth Middleware
	g := e.Group("/security")

	//added BasicAuth Middlewar
	g.Use(middleware.BasicAuth(middlewares.ValidateUser))

	//group with Basicauth sub path
	g.GET("/melta", handlers.MeltaGroup())

	//group with Basicauth sub path
	g.GET("/school", handlers.SchoolGroup())

	// middlewares
	e.Use(middlewares.MidCustomContext)

	// middleware cors
	//e.Use(middleware.CORS())

	//middlewares custom cors
	e.Use(middleware.CORSWithConfig(middlewares.CorsPolicy()))

	//request model validator.
	e.Validator = &CustomValidator{validator: validator.New()}

	//routes
	e.GET("/", handlers.Starts())

	//routes with create cookies
	e.GET("/cat", handlers.CreateCookies())

	//routes with read cookies
	e.GET("/dog", handlers.ReadCookies())

	//routes request binder post
	e.POST("/requ", handlers.RequestBinder())

	//routes request binder with validation
	e.POST("/vali", handlers.ValidatorsReq())

	//routess response
	e.GET("/respo", handlers.ResponseHandler())

	//routess html response
	e.GET("/index.html", handlers.HtmlResponse())

	//routess json pretty
	e.GET("/jsonreety", handlers.JsonPrettyResponse())

	//routess xml response
	e.GET("/xmlresponse", handlers.XmlResponse())

	//routing Queryparam
	e.GET("/param", handlers.QueryParam())

	//routing Param
	e.GET("/param/:name", handlers.ParamPath())

	//routing with db insert into
	e.POST("/mysql", handlers.DatabaseInsert(MysqlPool))

	//start server
	e.Logger.Fatal(e.Start(":4200"))
}

