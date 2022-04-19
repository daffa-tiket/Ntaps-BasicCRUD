package pegawai

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/tiketerror"
	"github.com/tiket/learn-crud/interfaces"
	"github.com/tiket/learn-crud/interfaces/pegawai"
	"github.com/tiket/learn-crud/shared"
	"github.com/tiket/learn-crud/shared/dto"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"

	validator_mock "github.com/tiket/learn-crud/mocks/validator"
)

func Test_CreatePegawai(t *testing.T) {
	var (
		ctrl              = gomock.NewController(t)
		toBaseResponseDto = func(reader io.Reader) *dto.BaseResponseDto {
			var (
				res dto.BaseResponseDto
			)
			_ = json.NewDecoder(reader).Decode(&res)
			return &res
		}
	)

	defer ctrl.Finish()

	t.Run("when failed to bind request dto", func(t *testing.T) {
		var (
			ec            = echo.New()
			req           = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{---"))
			rec           = httptest.NewRecorder()
			ctx           = ec.NewContext(req, rec)
			controller, _ = NewController(shared.Holder{}, interfaces.Holder{})
		)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		if assert.NoError(t, controller.CreatePegawai(ctx)) {
			res := toBaseResponseDto(rec.Body)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, tiketerror.BAD_REQUEST, res.Code)
		}
	})

	t.Run("when failed to validate request", func(t *testing.T) {
		var (
			ec            = echo.New()
			req           = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec           = httptest.NewRecorder()
			ctx           = ec.NewContext(req, rec)
			validator     = validator_mock.NewMockValidator(ctrl)
			controller, _ = NewController(shared.Holder{}, interfaces.Holder{})
		)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		validator.EXPECT().
			Validate(gomock.Any()).
			Return(errors.New("failed to validate"))

		ec.Validator = validator

		if assert.NoError(t, controller.CreatePegawai(ctx)) {
			res := toBaseResponseDto(rec.Body)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, tiketerror.BAD_REQUEST, res.Code)
			assert.Equal(t, "failed to validate", res.Errors[0])

		}
	})

	t.Run("when failed to call service", func(t *testing.T) {
		var (
			req       = httptest.NewRequest(http.MethodGet, "/", strings.NewReader("{}"))
			rec       = httptest.NewRecorder()
			ec        = echo.New()
			ctx       = ec.NewContext(req, rec)
			validator = validator_mock.NewMockValidator(ctrl)
			service   = pegawai.NewMockViewService(ctrl)
		)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		validator.EXPECT().
			Validate(gomock.Any()).
			Return(nil)

		ec.Validator = validator

		service.EXPECT().
			CreatePegawai(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("failed"))

		controller, _ := NewController(shared.Holder{}, interfaces.Holder{PegawaiViewService: service})

		if assert.NoError(t, controller.CreatePegawai(ctx)) {
			res := toBaseResponseDto(rec.Body)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, tiketerror.BAD_REQUEST, res.Code)
			assert.Equal(t, "failed", res.Errors[0])

		}
	})

	t.Run("when succesfully call service", func(t *testing.T) {
		var (
			req       = httptest.NewRequest(http.MethodGet, "/", strings.NewReader("{}"))
			rec       = httptest.NewRecorder()
			ec        = echo.New()
			ctx       = ec.NewContext(req, rec)
			validator = validator_mock.NewMockValidator(ctrl)
			service   = pegawai.NewMockViewService(ctrl)
		)

		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		validator.EXPECT().
			Validate(gomock.Any()).
			Return(nil)

		service.EXPECT().
			CreatePegawai(gomock.Any(), gomock.Any()).
			Return(&pegawai_dto.CreatePegawaiResponseDto{}, nil)

		ec.Validator = validator

		controller, _ := NewController(shared.Holder{}, interfaces.Holder{PegawaiViewService: service})

		if assert.NoError(t, controller.CreatePegawai(ctx)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
