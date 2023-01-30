package vacation

import (
	"github.com/mitchellh/mapstructure"
)

func wheneverVacationRegistered_Use(data map[string]interface{}){
	
	event := NewVacationRegistered()
	mapstructure.Decode(data,&event)

	Use(event);
}

func wheneverVacationCancelled_Add(data map[string]interface{}){
	
	event := NewVacationCancelled()
	mapstructure.Decode(data,&event)

	Add(event);
}

func wheneverVacationRejected_Add(data map[string]interface{}){
	
	event := NewVacationRejected()
	mapstructure.Decode(data,&event)

	Add(event);
}

func wheneverEmployeeJoined_RegisterUser(data map[string]interface{}){
	
	event := NewEmployeeJoined()
	mapstructure.Decode(data,&event)

	RegisterUser(event);
}

func wheneverVacationDaysInsufficient_Update(data map[string]interface{}){
	
	event := NewVacationDaysInsufficient()
	mapstructure.Decode(data,&event)

	Update(event);
}

