package backend

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"

	"google.golang.org/appengine"
)

func init() {
	e := echo.New()

	e.Use(AppContext)
	e.GET("/", handle)

	s := standard.New("")
	s.SetHandler(e)
	http.Handle("/", s)
}

func handle(c echo.Context) error {
	appID := appengine.AppID(c.StdContext())
	html := fmt.Sprintf("<html><body>Hello, World! %s</body></html>", appID)
	return c.HTML(200, html)
}

func AppContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.SetStdContext(appengine.WithContext(c.StdContext(), c.Request().(*standard.Request).Request))
		return next(c)
	}
}
