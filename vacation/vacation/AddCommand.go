package vacation

type AddCommand struct{
    DayCount int `json:"dayCount" type:"int"`
    Reason string `json:"reason" type:"string"`
}
