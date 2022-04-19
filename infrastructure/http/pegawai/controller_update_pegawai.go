package pegawai

import (
	"context"
	echo "github.com/labstack/echo/v4"
	"github.com/tiket/TIX-HOTEL-UTILITIES-GO/util/tiketerror"
	pegawai_dto "github.com/tiket/learn-crud/shared/dto/pegawai"
	"github.com/tiket/learn-crud/shared/util"
	"strconv"
)

// All godoc
// @Tags pegawai-controller
// @Summary Pegawai
// @Description Put all mandatory parameter
// @Param X-Store-Id header string true "TIKETCOM" default(TIKETCOM)
// @Param X-Channel-Id header string true "DESKTOP" default(DESKTOP)
// @Param X-Currency header string true "IDR" default(IDR)
// @Param X-Request-Id header string true "ee35b46d" default(ee35b46d)
// @Param X-Service-Id header string true "gateway" default(gateway)
// @Param X-Username header string true "guest" default(guest)
// @Param X-Account-Id header string true "1" default(1)
// @Param X-Reseller-Id header string true "1" default(1)
// @Param X-Identity header string true "02f0e6bf" default(02f0e6bf)
// @Param X-Business-Id header string true "1" default(1)
// @Param X-Login-Media header string true "GOOGLE" default(GOOGLE)
// @Param Accept-Language header string true "EN" default(EN)
// @Param X-Forwarded-Host header string true "192.168.1.1" default(192.168.1.1)
// @Param True-Client-Ip header string true "192.168.1.1" default(192.168.1.1)
// @Param UpdatePegawaiRequestDto body pegawai.UpdatePegawaiRequestDto true "UpdatePegawaiRequestDto"
//
// @Accept json
// @Produce json
// @Success 200 {object} pegawai.UpdatePegawaiResponseDto
// @Router /pegawai/{id} [put]
func (c *Controller) UpdatePegawai(ec echo.Context) error {
	var (
		ctx = context.Background()
		//mandatory = util.GetMandatoryRequestV2(ec)
		request pegawai_dto.UpdatePegawaiRequestDto
	)

	if err := ec.Bind(&request); err != nil {
		return util.Response(ec, nil, tiketerror.New(tiketerror.BAD_REQUEST, err))
	}

	//request.MandatoryRequestV2Dto = mandatory
	//if err := ec.Validate(&request); err != nil {
	//	return util.Response(ec, nil, tiketerror.New(tiketerror.BAD_REQUEST, err))
	//}

	int_conv, _ := strconv.Atoi(ec.Param("id"))
	request.ID = int_conv
	response, err := c.InterfacesHolder.PegawaiViewService.UpdatePegawai(ctx, request)

	if err != nil {
		return util.Response(ec, nil, tiketerror.New(tiketerror.BAD_REQUEST, err))
	}

	return util.Response(ec, response, nil)

}
