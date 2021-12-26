package models

type Citizen struct {
	Code       string `json:"code"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Gender     string `json:"gender"`

	DateOfBirth   string `json:"date_of_birth"`
	Age           int    `json:"age"`
	DateOfJoining string `json:"date_of_joining"`
	ReligionId    int    `json:"religion_id"`
	Avatar        string `json:"avatar"`

	CollaboratorName  string `json:"collaborator_name"`
	CollaboratorPhone string `json:"collaborator_phone"`

	TemporaryAddress string `json:"temporary_address"`
	Address          string `json:"address"`
	Major            string `json:"major"`
}
