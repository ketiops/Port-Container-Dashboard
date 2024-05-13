package controller

type EqInformation struct {
	ID          int    `json:"id"`
	InspEqNo    string `json:"inspEqNo"`
	InspAuto    string `json:"inspAuto"`
	InspName    string `json:"inspName,omitempty"`
	InspLoc     string `json:"inspLoc,omitempty"`
	InspContact string `json:"inspContact,omitempty"`
	QDate       string `json:"qDate"`
}

type EqState struct {
	ID           int    `json:"id"`
	InspEqNo     string `json:"inspEqNo"`
	InspEqStatus string `json:"inspEqStatus"`
	QDate        string `json:"qDate"`
}
