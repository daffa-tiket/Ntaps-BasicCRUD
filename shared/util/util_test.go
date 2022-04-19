package util

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/tiketerror"
	"github.com/tiket/learn-crud/shared/dto"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetMandatoryRequest(t *testing.T) {
	var (
		ec = echo.New()
	)

	t.Run("when username is guest", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		req.Header.Set("username", "guest")
		res := GetMandatoryRequest(ctx)
		assert.Equal(t, 0, res.Login)
	})

	t.Run("when username is not guest", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		req.Header.Set("username", "admin")
		res := GetMandatoryRequest(ctx)
		assert.Equal(t, 1, res.Login)
	})

	t.Run("when lang is not set", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		req.Header.Set("lang", "")
		res := GetMandatoryRequest(ctx)
		assert.Equal(t, "en", res.Language)
	})

	t.Run("when lang is set", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		req.Header.Set("lang", "id")
		res := GetMandatoryRequest(ctx)
		assert.Equal(t, "id", res.Language)
	})

	t.Run("when currency is not set", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		req.Header.Set("currency", "")
		res := GetMandatoryRequest(ctx)
		assert.Equal(t, "IDR", res.Currency)
	})

	t.Run("when currency is set", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		req.Header.Set("currency", "USD")
		res := GetMandatoryRequest(ctx)
		assert.Equal(t, "USD", res.Currency)
	})

	t.Run("when true-client-ip is not set", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		req.Header.Set("true-client-ip", "")
		res := GetMandatoryRequest(ctx)
		assert.Equal(t, "127.0.0.1", res.CustomerIPAddress)
	})

	t.Run("when true-client-ip is set", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		req.Header.Set("true-client-ip", "192.168.0.1")
		res := GetMandatoryRequest(ctx)
		assert.Equal(t, "192.168.0.1", res.CustomerIPAddress)
	})

	t.Run("default value", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		req.Header.Set("username", "admin")
		req.Header.Set("lang", "id")
		req.Header.Set("currency", "IDR")
		req.Header.Set("true-client-ip", "192.168.0.1")
		req.Header.Set("channelId", "some-channel")
		req.Header.Set("customerSessionId", "some-session")
		req.Header.Set("requestId", "some-request")
		req.Header.Set("serviceId", "some-service")
		req.Header.Set("storeId", "some-store")
		req.Header.Set("user-agent", "some-agent")
		req.Header.Set("resellerId", "some-reseller")

		res := GetMandatoryRequest(ctx)

		assert.Equal(t, "admin", res.Username)
		assert.Equal(t, "id", res.Language)
		assert.Equal(t, "IDR", res.Currency)
		assert.Equal(t, "192.168.0.1", res.CustomerIPAddress)
		assert.Equal(t, "some-channel", res.ChannelID)
		assert.Equal(t, "some-session", res.CustomerSessionId)
		assert.Equal(t, "some-request", res.RequestID)
		assert.Equal(t, "some-service", res.ServiceID)
		assert.Equal(t, "some-store", res.StoreID)
		assert.Equal(t, "some-agent", res.CustomerUserAgent)
		assert.Equal(t, "some-reseller", res.ResellerID)
	})
}

func TestGetMandatoryRequestV2(t *testing.T) {
	var (
		ec  = echo.New()
		req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
		rec = httptest.NewRecorder()
		ctx = ec.NewContext(req, rec)
	)

	req.Header.Set("X-Store-Id", "some-store-id")
	req.Header.Set("X-Channel-Id", "some-channel-id")
	req.Header.Set("X-Request-Id", "some-request-id")
	req.Header.Set("X-Identity", "some-identity")
	req.Header.Set("X-Account-Id", "some-account-id")
	req.Header.Set("X-Username", "some-username")
	req.Header.Set("Accept-Language", "some-language")
	req.Header.Set("X-Currency", "some-currency")
	req.Header.Set("X-Reseller-Id", "")
	req.Header.Set("X-Business-Id", "")
	req.Header.Set("X-Login-Media", "")
	req.Header.Set("X-Forwarded-For", "")
	req.Header.Set("True-Client-Ip", "")

	res := GetMandatoryRequestV2(ctx)
	assert.Equal(t, "some-store-id", res.StoreID)
	assert.Equal(t, "some-channel-id", res.ChannelID)
	assert.Equal(t, "some-request-id", res.RequestID)
	assert.Equal(t, "some-identity", res.Identity)
	assert.Equal(t, "some-account-id", res.AccountID)
	assert.Equal(t, "some-username", res.Username)
	assert.Equal(t, "some-language", res.Language)
	assert.Equal(t, "some-currency", res.Currency)
	assert.Equal(t, "1", res.ResellerID)
	assert.Equal(t, "1", res.BusinessID)
	assert.Equal(t, "GOOGLE", res.LoginMedia)
	assert.Equal(t, "127.0.0.1", res.ForwardedFor)
	assert.Equal(t, "127.0.0.1", res.TrueClientIP)
}

func TestResponse(t *testing.T) {
	var (
		ec = echo.New()
	)

	t.Run("when error is nil", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		assert.NoError(t, Response(ctx, nil, nil))
	})

	t.Run("when error is ErrorStandard", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		err := Response(ctx, nil, tiketerror.New("BAD_REQUEST", errors.New("some error")))
		assert.NoError(t, err)

		base := dto.BaseResponseDto{}
		json.Unmarshal(rec.Body.Bytes(), &base)
		assert.Equal(t, "BAD_REQUEST", base.Code)
		assert.Equal(t, "Bad request", base.Message)
	})

	t.Run("when error is not ErrorStandard", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
			rec = httptest.NewRecorder()
			ctx = ec.NewContext(req, rec)
		)

		err := Response(ctx, nil, errors.New("some error"))
		assert.NoError(t, err)

		base := dto.BaseResponseDto{}
		json.Unmarshal(rec.Body.Bytes(), &base)
		assert.Equal(t, "SYSTEM_ERROR", base.Code)
		assert.Equal(t, "Contact our team", base.Message)
	})

}

func Test_getHeaderV2OrDefault(t *testing.T) {
	t.Run("when value is not found", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
		)

		assert.Equal(t, getHeaderV2OrDefault(req.Header, "some-missing-header", "default-value"), "default-value")

	})

	t.Run("when value is found", func(t *testing.T) {
		var (
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
		)

		req.Header.Set("some-header", "some-value")
		assert.Equal(t, getHeaderV2OrDefault(req.Header, "some-header", "not-set"), "some-value")
	})
}
