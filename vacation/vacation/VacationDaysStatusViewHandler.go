package vacation

import (
	"github.com/mitchellh/mapstructure"
	"log"
)

func whenVacationDaysIntialized_then_CREATE_1 (inputEvent map[string]interface{}) {

	vacationDaysIntialized := NewVacationDaysIntialized()
	mapstructure.Decode(inputEvent, &vacationDaysIntialized)

	vacationDaysStatus := &VacationDaysStatus{}
	vacationDaysStatus.UserId = vacationDaysIntialized.UserId
	vacationDaysStatus.DaysLeft = vacationDaysIntialized.DayCount

	// view 레파지 토리에 save
	repository := VacationDaysStatusRepository()
	err := repository.save(vacationDaysStatus)
	if err != nil {
		// TODO error control
		log.Printf("Create error: %v \n", err)
	}
}

func whenVacationDaysAdded_then_UPDATE_1 (inputEvent map[string]interface{}) {

	vacationDaysAdded := NewVacationDaysAdded()
	mapstructure.Decode(inputEvent,&vacationDaysAdded)

	var vacationDaysStatuss []VacationDaysStatus
	repository := VacationDaysStatusRepository()
	// FIXME geom lib define snake_case as column name (eg: user_id), so if your query column is 'userId' then change 'user_id'
	err := repository.db.Where("userId = ?", vacationDaysAdded.UserId).Find(&vacationDaysStatuss).Error
	if err != nil {
		// TODO error control
		log.Printf("Select error: %v \n", err)
	}
	for _, viewEntity := range vacationDaysStatuss {
		viewEntity.DaysLeft = vacationDaysAdded.DayCount
		err1 := repository.db.Updates(viewEntity).Error
		if err1 != nil {
			// TODO error control
			log.Printf("Update error: %v \n", err1)
		}
	}
}
func whenVacationDaysUsed_then_UPDATE_2 (inputEvent map[string]interface{}) {

	vacationDaysUsed := NewVacationDaysUsed()
	mapstructure.Decode(inputEvent,&vacationDaysUsed)

	var vacationDaysStatuss []VacationDaysStatus
	repository := VacationDaysStatusRepository()
	// FIXME geom lib define snake_case as column name (eg: user_id), so if your query column is 'userId' then change 'user_id'
	err := repository.db.Where("userId = ?", vacationDaysUsed.UserId).Find(&vacationDaysStatuss).Error
	if err != nil {
		// TODO error control
		log.Printf("Select error: %v \n", err)
	}
	for _, viewEntity := range vacationDaysStatuss {
		viewEntity.DaysLeft = vacationDaysUsed.DayCount
		err1 := repository.db.Updates(viewEntity).Error
		if err1 != nil {
			// TODO error control
			log.Printf("Update error: %v \n", err1)
		}
	}
}

