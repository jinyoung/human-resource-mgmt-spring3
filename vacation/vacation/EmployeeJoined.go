package vacation

import (
	"time"
)

type EmployeeJoined struct{
	EventType string	`json:"eventType" type:"string"`
	TimeStamp string 	`json:"timeStamp" type:"string"`
	UserId string `json:"userId" type:"string"` 
	Name string `json:"name" type:"string"` 
	Email string `json:"email" type:"string"` 
	
}

func NewEmployeeJoined() *EmployeeJoined{
	event := &EmployeeJoined{EventType:"EmployeeJoined", TimeStamp:time.Now().String()}

	return event
}
