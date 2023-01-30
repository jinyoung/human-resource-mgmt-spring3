package vacation

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
	"fmt"
	"vacation/external"
)

type Vacation struct {
	gorm.Model
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	StartDate undefined `json:"startDate"`
	EndDate undefined `json:"endDate"`
	Reason string `json:"reason"`
	UserId string `json:"userId"`
	Days int `json:"days"`
	Status string `json:"status"`

}

func (self *Vacation) onPostPersist() (err error){
	vacationRegistered := NewVacationRegistered()
	model.Copy(vacationRegistered, self)

	Publish(vacationRegistered)

	// Get request from Vacation
	vacation := &external.Vacation{}
	resp := external.GetVacation(vacation.Id)
	fmt.Println(resp)

	return nil
}
func (self *Vacation) onPrePersist() (err error){ return nil }
func (self *Vacation) onPreUpdate() (err error){ return nil }
func (self *Vacation) onPostUpdate() (err error){ return nil }
func (self *Vacation) onPreRemove() (err error){ return nil }
func (self *Vacation) onPostRemove() (err error){ return nil }

func (self *Vacation) Cancel(){
	vacationCancelled := NewVacationCancelled()
	model.Copy(vacationCancelled, self)
	Publish(vacationCancelled)
}
func (self *Vacation) Approve(approveCommand *{namePascalCase}}Command){
	vacationApproved := NewVacationApproved()
	model.Copy(vacationApproved, self)
	Publish(vacationApproved)
}
func (self *Vacation) ConfirmUsed(){
	vacationUsed := NewVacationUsed()
	model.Copy(vacationUsed, self)
	Publish(vacationUsed)
}
func (self *Vacation) Update(){
	vacationRejected := NewVacationRejected()
	model.Copy(vacationRejected, self)
	Publish(vacationRejected)
}

