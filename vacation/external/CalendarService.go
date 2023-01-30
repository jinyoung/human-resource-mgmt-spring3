package external 

import (
	"github.com/go-resty/resty/v2"
	"fmt"
	"vacation/config"
)

var client = resty.New()

func GetCalendar( userId string) *resty.Response{
	options := config.Reader(config.GetMode())
	target := fmt.Sprintf("https://%s/%s/%s", options["api_url_schedule"], "calendars" , userId )
	resp, _ := client.R().Get(target)

	return resp
}
