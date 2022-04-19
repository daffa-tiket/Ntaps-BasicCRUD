package util

import (
	"net/http"
	"strings"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/tiketerror"
	"github.com/tiket/learn-crud/shared/dto"
)

// Deprecated: use GetMandatoryRequestV2
func GetMandatoryRequest(ec echo.Context) dto.MandatoryRequestDto {
	lang := ec.Request().Header.Get("lang")
	username := ec.Request().Header.Get("username")
	currency := ec.Request().Header.Get("currency")
	clientIP := ec.Request().Header.Get("true-client-ip")

	login := 0
	if strings.ToLower(username) != "guest" {
		login = 1
	}
	if lang == "" {
		lang = "en"
	}
	if currency == "" {
		currency = "IDR"
	}
	if clientIP == "" {
		clientIP = "127.0.0.1"
	}

	mandatory := dto.MandatoryRequestDto{
		ChannelID:         ec.Request().Header.Get("channelId"),
		CustomerSessionId: ec.Request().Header.Get("customerSessionId"),
		RequestID:         ec.Request().Header.Get("requestId"),
		ServiceID:         ec.Request().Header.Get("serviceId"),
		StoreID:           ec.Request().Header.Get("storeId"),
		CustomerUserAgent: ec.Request().Header.Get("user-agent"),
		ResellerID:        ec.Request().Header.Get("resellerId"),
		CustomerIPAddress: clientIP,
		Language:          lang,
		Currency:          currency,
		Username:          username,
		Login:             login,
	}

	return mandatory
}

func GetMandatoryRequestV2(ec echo.Context) dto.MandatoryRequestV2Dto {
	var (
		header = ec.Request().Header
	)

	return dto.MandatoryRequestV2Dto{
		StoreID:      getHeaderV2OrDefault(header, "X-Store-Id", ""),
		ChannelID:    getHeaderV2OrDefault(header, "X-Channel-Id", ""),
		RequestID:    getHeaderV2OrDefault(header, "X-Request-Id", ""),
		ServiceID:    getHeaderV2OrDefault(header, "X-Service-Id", ""),
		Identity:     getHeaderV2OrDefault(header, "X-Identity", ""),
		AccountID:    getHeaderV2OrDefault(header, "X-Account-Id", ""),
		Username:     getHeaderV2OrDefault(header, "X-Username", ""),
		Language:     getHeaderV2OrDefault(header, "Accept-Language", ""),
		Currency:     getHeaderV2OrDefault(header, "X-Currency", ""),
		ResellerID:   getHeaderV2OrDefault(header, "X-Reseller-Id", "1"),
		BusinessID:   getHeaderV2OrDefault(header, "X-Business-Id", "1"),
		LoginMedia:   getHeaderV2OrDefault(header, "X-Login-Media", "GOOGLE"),
		ForwardedFor: getHeaderV2OrDefault(header, "X-Forwarded-For", "127.0.0.1"),
		TrueClientIP: getHeaderV2OrDefault(header, "True-Client-Ip", "127.0.0.1"),
	}
}

func Response(ec echo.Context, data interface{}, err error) error {
	if err == nil {
		successResponse := tiketerror.New(tiketerror.SUCCESS, nil)
		return ec.JSON(successResponse.GetHTTPStatus(), &dto.BaseResponseDto{
			Code:       successResponse.GetCode(),
			Message:    successResponse.GetMessage(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	}
	if s, ok := err.(tiketerror.ErrorStandard); ok {
		return ec.JSON(s.GetHTTPStatus(), &dto.BaseResponseDto{
			Code:       s.GetCode(),
			Message:    s.GetMessage(),
			Errors:     s.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	} else {
		errResponse := tiketerror.New(tiketerror.SYSTEM_ERROR, err)
		return ec.JSON(errResponse.GetHTTPStatus(), &dto.BaseResponseDto{
			Code:       errResponse.GetCode(),
			Message:    errResponse.GetMessage(),
			Errors:     errResponse.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	}
}

func getHeaderV2OrDefault(header http.Header, key, value string) string {
	var (
		val = header.Get(key)
	)

	if len(val) == 0 {
		return value
	}

	return val
}
