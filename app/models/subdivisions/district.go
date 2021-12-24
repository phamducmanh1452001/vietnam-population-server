package models

type District struct {
	Name       string           `json:"name"`
	Code       string           `json:"code"`
	SuperCode  string           `json:"super_code"`
	Population uint32           `json:"population"`
	Level      SubdivisionLevel `json:"level"`
	Area       float64          `json:"area"`
}
