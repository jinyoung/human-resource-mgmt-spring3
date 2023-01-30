package vacation

type VacationDaysStatus struct {
	UserId int `gorm:"primaryKey" json:"id" type:"int"`
	DaysLeft int `json:"daysLeft" type:"int"`
}