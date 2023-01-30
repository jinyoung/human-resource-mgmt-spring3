package vacation

import (
	"time"
)

type VacationDaysInsufficient struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	Id int `json:"id" type:"int"` 
	VacationId string `json:"vacationId" type:"string"` 
	
}

func NewVacationDaysInsufficient() *VacationDaysInsufficient{
	event := &VacationDaysInsufficient{EventType:"VacationDaysInsufficient", TimeStamp:time.Now().String()}

	return event
}
