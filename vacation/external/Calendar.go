package external

type Calendar struct {
	UserId int `gorm:"primaryKey" json:"id" type:"int"`
	Events  `json:"events" type:""`
}

func (self *Calendar) getPrimaryKey() int {
	// FIXME if PrimaryKey is multi value, than change this method
	return self.UserId
}
