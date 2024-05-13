package controller

type PreInformation struct {
	ID       int    `json:"id"`
	InspEqNo string `json:"inspEqNo"`
	CntrNo   string `json:"cntrNo"`
	TruckNo  string `json:"truckNo"`
	TypeNo   string `json:"typeNo"`
	QDate    string `json:"qDate"`
}

type ContainerSpec struct {
	ID             int    `json:"id"`
	InspEqNo       string `json:"inspEqNo"`
	InspNo         string `json:"inspNo"`
	InspStartTime  string `json:"inspStartTime,omitempty"`
	InspEndTime    string `json:"inspEndTime,omitempty"`
	PckMatch       string `json:"pckMatch"`
	InspRsltCD     string `json:"inspRsltCD"`
	DetectionCnt   string `json:"detectionCnt"`
	FaultCD        string `json:"faultCD"`
	InspRsltImgDir string `json:"inspRsltImgDir,omitempty"`
	QDate          string `json:"qDate"`
}

type Remark struct {
	RemarkID      int    `json:"remarkId"`
	InspRemark    string `json:"inspRemark"`
	InformationID int    `json:"informationId"`
	QDate         string `json:"qDate"`
}
