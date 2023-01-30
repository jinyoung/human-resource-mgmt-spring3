package main

import(
	"log"
	"vacation/vacation"
	"vacation/config"
)

func main() {
	config.ConfigMode()
	options := config.Reader(config.GetMode())
	log.Printf("option : %v", options)
	vacation.VacationDBInit()
	vacation.VacationDaysLeftDBInit()
	vacation.VacationDaysStatusDBInit()
	vacation.VacationStatusDBInit()
	vacation.VacationDaysLeftDBInit()
	go vacation.InitProducer(config.GetMode())
	// view 와 같이 사용시 InitConsumer 가 중복으로 호출될수 있으니, 하나는 삭제 필요
	go vacation.InitConsumer(config.GetMode())
	// policy 와 같이 사용시 InitConsumer 가 중복으로 호출될수 있으니, 하나는 삭제 필요
	go vacation.InitConsumer(config.GetMode())
	// policy 와 같이 사용시 InitConsumer 가 중복으로 호출될수 있으니, 하나는 삭제 필요
	go vacation.InitConsumer(config.GetMode())
	// policy 와 같이 사용시 InitConsumer 가 중복으로 호출될수 있으니, 하나는 삭제 필요
	go vacation.InitConsumer(config.GetMode())
	e := vacation.RouteInit()

	e.Logger.Fatal(e.Start(":8081"))
}
