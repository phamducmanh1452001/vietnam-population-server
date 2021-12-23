package models

type Citizen struct {
	Code          string `json:"code"`
	FirstName     string `json:"first_name"`
	MiddleName    string `json:"middle_name"`
	LastName      string `json:"last_name"`
	Gender        string `json:"gender"`
	DateOfBirth   string `json:"date_of_birth"`
	Age           int    `json:"age"`
	Weight        int    `json:"weight"`
	DateOfJoining string `json:"date_of_joining"`
	Religion      string `json:"religion"`
}
