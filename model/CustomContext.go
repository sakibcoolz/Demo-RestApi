package model

import "github.com/labstack/echo"

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}
