package vacation

import (
	"time"
)

type VacationCancelled struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id string `json:"id" type:"string"` 
	StartDate  `json:"startDate" type:""` 
	EndDate  `json:"endDate" type:""` 
	Reason string `json:"reason" type:"string"` 
	UserId string `json:"userId" type:"string"` 
	Days int `json:"days" type:"int"` 
	
}

func NewVacationCancelled() *VacationCancelled{
	event := &VacationCancelled{EventType:"VacationCancelled", TimeStamp:time.Now().String()}

	return event
}
