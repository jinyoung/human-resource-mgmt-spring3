package external

type Vacation struct {
	Id int `gorm:"primaryKey" json:"id" type:"int"`
	StartDate  `json:"startDate" type:""`
	EndDate  `json:"endDate" type:""`
	Reason string `json:"reason" type:"string"`
	UserId string `json:"userId" type:"string"`
	Days int `json:"days" type:"int"`
	Status string `json:"status" type:"string"`
}

func (self *Vacation) getPrimaryKey() int {
	// FIXME if PrimaryKey is multi value, than change this method
	return self.Id
}
