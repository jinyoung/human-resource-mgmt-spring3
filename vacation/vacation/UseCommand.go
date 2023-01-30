package vacation

type UseCommand struct{
    DayCount int `json:"dayCount" type:"int"`
    Reason string `json:"reason" type:"string"`
    VacationId string `json:"vacationId" type:"string"`
}
