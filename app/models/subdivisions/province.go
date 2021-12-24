package models

type SubdivisionLevel struct {
	Name   string `json:"name"`
	Number uint8  `json:"number"`
}

type Province struct {
	Name       string           `json:"name"`
	Code       string           `json:"code"`
	Population uint32           `json:"population"`
	Level      SubdivisionLevel `json:"level"`
	Area       float64          `json:"area"`
}
