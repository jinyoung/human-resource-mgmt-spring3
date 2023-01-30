package vacation

import (
	"time"
)

type VacationDaysUsed struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	UserId string `json:"userId" type:"string"` 
	DayCount int `json:"dayCount" type:"int"` 
	Reason string `json:"reason" type:"string"` 
	
}

func NewVacationDaysUsed() *VacationDaysUsed{
	event := &VacationDaysUsed{EventType:"VacationDaysUsed", TimeStamp:time.Now().String()}

	return event
}
