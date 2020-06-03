package uyeg

type RemapFormatV1 struct {
	Version          uint16     `json:"ver"`
	GatewayID        string     `json:"gateway"`
	MacID            string     `json:"mac"`
	Time             string     `json:"time"`
	Temp             float64    `json:"Temp"`
	Humid            float64    `json:"Humid"`
	ActiveConsum     float64    `json:"ActiveConsum"`
	ReactiveConsum   float64    `json:"ReactiveConsum"`
	Power            float64    `json:"Power"`
	TotalRunningHour float64    `json:"TotalRunningHour"`
	MCCounter        float64    `json:"MCCounter"`
	PT100            float64    `json:"PT100"`
	FaultNumber      float64    `json:"FaultNumber"`
	FaultRST         float64    `json:"FaultRST"`
	Values           []Depth2V1 `json:"Values"`
}

type RemapFormatV2 struct {
	Version          uint16     `json:"ver"`
	GatewayID        string     `json:"gateway"`
	MacID            string     `json:"mac"`
	Time             string     `json:"time"`
	Temp             float64    `json:"Temp"`
	Humid            float64    `json:"Humid"`
	ReactivePower    float64    `json:"ReactivePower"`
	ActiveConsum     float64    `json:"ActiveConsum"`
	ReactiveConsum   float64    `json:"ReactiveConsum"`
	Power            float64    `json:"Power"`
	RunningHour      float64    `json:"RunningHour"`
	TotalRunningHour float64    `json:"TotalRunningHour"`
	MCCounter        float64    `json:"MCCounter"`
	PT100            float64    `json:"PT100"`
	FaultNumber      float64    `json:"FaultNumber"`
	OverCurrR        float64    `json:"OverCurrR"`
	OverCurrS        float64    `json:"OverCurrS"`
	OverCurrT        float64    `json:"OverCurrT"`
	FaultRST         float64    `json:"FaultRST"`
	Values           []Depth2V1 `json:"Values"`
}

type Depth2V1 struct {
	Time        string  `json:"time"`
	Status      bool    `json:"status"`
	Curr        float64 `json:"Curr"`
	CurrR       float64 `json:"CurrR"`
	CurrS       float64 `json:"CurrS"`
	CurrT       float64 `json:"CurrT"`
	Volt        float64 `json:"Volt"`
	VoltR       float64 `json:"VoltR"`
	VoltS       float64 `json:"VoltS"`
	VoltT       float64 `json:"VoltT"`
	ActivePower float64 `json:"ActivePower"`
	Ground      float64 `json:"Ground"`
	V420        float64 `json:"420"`
}
