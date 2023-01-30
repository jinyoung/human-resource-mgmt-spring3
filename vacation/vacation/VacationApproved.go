package vacation

import (
	"time"
)

type VacationApproved struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id string `json:"id" type:"string"` 
	StartDate  `json:"startDate" type:""` 
	EndDate  `json:"endDate" type:""` 
	Reason string `json:"reason" type:"string"` 
	
}

func NewVacationApproved() *VacationApproved{
	event := &VacationApproved{EventType:"VacationApproved", TimeStamp:time.Now().String()}

	return event
}
