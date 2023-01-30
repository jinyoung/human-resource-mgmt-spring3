package vacation

import (
	"gopkg.in/jeevatkm/go-model.v1"
	
	"gorm.io/gorm"
	"fmt"
	"vacation/external"
)

type VacationDaysLeft struct {
	gorm.Model
	UserId int `gorm:"primaryKey" json:"id" type:"int"`
	DayCount int `json:"dayCount"`

}

func (self *VacationDaysLeft) onPostPersist() (err error){
	vacationDaysIntialized := NewVacationDaysIntialized()
	model.Copy(vacationDaysIntialized, self)

	Publish(vacationDaysIntialized)
	vacationDaysInsufficient := NewVacationDaysInsufficient()
	model.Copy(vacationDaysInsufficient, self)

	Publish(vacationDaysInsufficient)

	return nil
}
func (self *VacationDaysLeft) onPrePersist() (err error){ return nil }
func (self *VacationDaysLeft) onPreUpdate() (err error){ return nil }
func (self *VacationDaysLeft) onPostUpdate() (err error){ return nil }
func (self *VacationDaysLeft) onPreRemove() (err error){ return nil }
func (self *VacationDaysLeft) onPostRemove() (err error){ return nil }

func (self *VacationDaysLeft) Add(addCommand *{namePascalCase}}Command){
	vacationDaysAdded := NewVacationDaysAdded()
	model.Copy(vacationDaysAdded, self)
	Publish(vacationDaysAdded)
}
func (self *VacationDaysLeft) Use(useCommand *{namePascalCase}}Command){
	vacationDaysUsed := NewVacationDaysUsed()
	model.Copy(vacationDaysUsed, self)
	Publish(vacationDaysUsed)
}

