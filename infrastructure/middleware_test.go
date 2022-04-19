package infrastructure

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	mock_logger "github.com/tiket/learn-crud/mocks/logger"
	"github.com/tiket/learn-crud/shared"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Register(t *testing.T) {
	RegisterMiddleware(&Holder{
		SharedHolder: shared.Holder{
			Echo: echo.New(),
		},
	})
}

func Test_HelloWorldMiddleware(t *testing.T) {
	var (
		ctrl = gomock.NewController(t)
	)

	defer ctrl.Finish()

	t.Run("hello world middleware called", func(t *testing.T) {
		var (
			logger = mock_logger.NewMockLogger(ctrl)
			sh     = shared.Holder{Logger: logger}
			ec     = echo.New()
			req    = httptest.NewRequest(http.MethodGet, "/", strings.NewReader("{}"))
			rec    = httptest.NewRecorder()
		)
		logger.EXPECT().Info("hello world")
		ec.Use(HelloWorldMiddleware(&Holder{SharedHolder: sh}))

		ec.GET("/", func(c echo.Context) error {
			return c.String(200, "ok")
		})

		ec.ServeHTTP(rec, req)
	})
}
