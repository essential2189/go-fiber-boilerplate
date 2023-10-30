package controller

import (
	"go-boilerplate/app"
	"go-boilerplate/app/core/consts"
	"go-boilerplate/app/core/helper/logger"
	"go-boilerplate/app/core/helper/resty"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Table() []app.Mapping
	Temp(c *fiber.Ctx) error
}

type controller struct {
	resty *resty.HttpClient
}

func NewController(resty *resty.HttpClient) Controller {
	return controller{
		resty: resty,
	}
}

func (ctrl controller) Table() []app.Mapping {
	return []app.Mapping{
		{fiber.MethodGet, "/temp", ctrl.Temp},
	}
}

var log = logger.Get()

func (ctrl controller) Temp(c *fiber.Ctx) error {
	requestInfo := resty.RequestInfo{
		Uri:          "http://wavve-apis/platform/accesscheck?apikey=E5F3E0D30947AA5440556471321BB6D9&device=none&partner=none&pooqzone=none&region=none&drm=none&targetage=none'",
		Method:       resty.MethodGET,
		Headers:      map[string]string{"Content-Type": "application/json", consts.WavveCredentialHeader: "SiOCXq35eH2tsR2Mp7479lt/RBia2ksDbYCmgl/I1MDYd9GqrKePTy1sWC8wx4Y1tKJcodH7e6eIOscEO+M5TNZcaf3HtokBhBGoTYF1Voq9oA463XGJriAt370dIfXpDWc18Go2NXIIX8D/pzI+uoWFb6rwNpCZ7cYrF3qikmVB5TLxbt3QTfp7SqZoIbRoKNE9rH3VEwcbXJPnckTeJqs1QVV6xtoanUmDB53YVwFUmoy9F3L4FnE3ywWoMfrfrt8FVWzK/+dpFareWXgqKyeKBz42rdH8Dviy30dEXRexRG/GwlG3jl34YMO2eGL0w7kl6m2+8juR6ameIaKImw=="},
		Query:        map[string]string{},
		Body:         nil,
		Timeout:      0,
		RetryCount:   0,
		RetryBackOff: 0,
		IsSkipSSL:    false,
	}
	res, _ := ctrl.resty.Request(requestInfo)

	log.Info("res: " + strconv.Itoa(res.StatusCode))
	return c.SendStatus(http.StatusOK)
}
