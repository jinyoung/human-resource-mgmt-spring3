package external 

import (
	"github.com/go-resty/resty/v2"
	"fmt"
	"vacation/config"
)

var client = resty.New()

func registerVacation(vacation Vacation) *resty.Response {
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s", options["api_url_vacation"], "vacations" )
	resp, _ := client.R().SetBody(vacation).Post(target)

	return resp



}
