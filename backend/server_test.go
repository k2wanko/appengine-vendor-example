package backend

import (
	"testing"

	"google.golang.org/appengine/aetest"

	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
)

func TestHandle(t *testing.T) {
	t.Parallel()
	e := echo.New()
	req := test.NewRequest(echo.GET, "/", nil)
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)

	err := appTestContext(handle)(c)

	if err != nil {
		t.Fatal(err)
	}

	if s := rec.Status(); s != 200 {
		t.Errorf("Status not ok: status = %d", s)
	}
}

func appTestContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, close, err := aetest.NewContext()
		if err != nil {
			return err
		}
		defer close()
		c.SetStdContext(ctx)
		return next(c)
	}
}
