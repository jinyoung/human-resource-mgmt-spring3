package vacation

import (
	"time"
)

type VacationDaysIntialized struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	UserId string `json:"userId" type:"string"` 
	DayCount int `json:"dayCount" type:"int"` 
	
}

func NewVacationDaysIntialized() *VacationDaysIntialized{
	event := &VacationDaysIntialized{EventType:"VacationDaysIntialized", TimeStamp:time.Now().String()}

	return event
}
