package handlers

import (
	"Demo-RestApi/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

//routes request response context
func Starts() echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		//cc := c.(*model.CustomContext)
		//cc.Bar()
		//cc.Foo()
		return c.String(http.StatusOK, "Hello, World!")
	}
}

func CreateCookies() echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		cookie:=new(http.Cookie)
		cookie.Name  = "sakib"
		cookie.Value = "mulla"
		cookie.Expires= time.Now().Add(24*time.Minute)
		c.SetCookie(cookie)
		return c.String(http.StatusOK, "create cookies")
	}
}

func ReadCookies() echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		cookies, err := c.Cookie("sakib")
		if err != nil {
			return err
		}
		fmt.Println(cookies.Name)
		fmt.Println(cookies.Value)
		return c.String(http.StatusOK, "read cookies")
	}
}

func RequestBinder() echo.HandlerFunc  {
	return func (c echo.Context) (err error) {
		u := new(model.User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	}
}

func DatabaseInsert(db *sql.DB) echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		u := new(model.User)
		if err := c.Bind(u); err != nil {
			return err
		}
		model.UserInsert(u, db)
		return c.JSON(http.StatusOK, u)
	}
}


func ValidatorsReq() echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		u:=new(model.User)
		if err = c.Bind(u); err != nil {
			return err
		}
		if err = c.Validate(u); err != nil {
			return c.JSON(http.StatusOK, model.Msg{Message: "Validation", Flags: "Error", Error: "error"})
		}
		return c.JSON(http.StatusOK, u)
	}
}

func ResponseHandler() echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		u := model.User{
			Email: "sakib.mulla@gmail.com",
			Name: "sakib",
		}
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusOK)
		return json.NewEncoder(c.Response()).Encode(u)
	}
}

func HtmlResponse() echo.HandlerFunc {
	return func (c echo.Context) (err error) {
		return c.HTML(http.StatusOK, "<marquee>Sakib Mulla</marquee>")
	}
}

func JsonPrettyResponse() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &model.User{
			Name:  "Jon",
			Email: "joe@labstack.com",
		}
		return c.JSONPretty(http.StatusOK, u, "  ")
	}
}


func XmlResponse() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := &model.User{
			Name:  "Jon",
			Email: "joe@labstack.com",
		}
		return c.XML(http.StatusOK, u)
	}
}


func QueryParam() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.String(http.StatusOK, "Query Param "+name)
	}
}

func ParamPath() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, "Param name : "+name)
	}
}


func MeltaGroup() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.QueryParam("name")
		return c.String(http.StatusOK, "Param name : "+name)
	}
}

func SchoolGroup() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Okay you are authenticated user ")
	}
}