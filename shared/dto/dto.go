package dto

type (
	MandatoryRequestDto struct {
		StoreID           string `json:"storeId" validate:"required"`
		ChannelID         string `json:"channelId" validate:"required"`
		RequestID         string `json:"requestId" validate:"required"`
		ServiceID         string `json:"serviceId" validate:"required"`
		Username          string `json:"username" validate:"required"`
		Language          string `json:"lang,omitempty"`
		Login             int    `json:"login,omitempty"`
		CustomerUserAgent string `json:"customerUserAgent,omitempty"`
		CustomerIPAddress string `json:"customerIpAddress,omitempty"`
		CustomerSessionId string `json:"customerSessionId,omitempty"`
		Currency          string `json:"currency,omitempty"`
		ResellerID        string `json:"resellerId,omitempty"`
		Version           string `json:"version,omitempty"`
		UserAgent         string `json:"user-agent,omitempty"`
		XForwardedHost    string `json:"x-forwarded-host,omitempty"`
	}

	MandatoryRequestV2Dto struct {
		StoreID      string `json:"X-Store-Id" validate:"required"`
		ChannelID    string `json:"X-Channel-Id" validate:"required"`
		RequestID    string `json:"X-Request-Id" validate:"required"`
		ServiceID    string `json:"X-Service-Id" validate:"required"`
		AccountID    string `json:"X-Account-Id" validate:"required"`
		Username     string `json:"X-Username" validate:"required"`
		Currency     string `json:"X-Currency" validate:"required"`
		Identity     string `json:"X-Identity"`
		ResellerID   string `json:"X-Reseller-Id"`
		BusinessID   string `json:"X-Business-Id"`
		LoginMedia   string `json:"X-Login-Media"`
		ForwardedFor string `json:"X-Forwarded-For"`
		TrueClientIP string `json:"True-Client-Ip"`
		Language     string `json:"Accept-Language" validate:"required"`
	}

	BaseResponseDto struct {
		Code       string      `json:"code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
		Errors     []string    `json:"errors"`
		ServerTime int64       `json:"serverTime"`
	}
)
