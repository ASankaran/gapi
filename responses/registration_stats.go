package responses

import (
	"../models"
)

type RegistrationStats struct {
	School                   []models.Stat      `json:"school,omitempty"`
	Transportation           []models.Stat      `json:"transportation,omitempty"`
	Diet                     []models.Stat      `json:"diet,omitempty"`
	ShirtSize                []models.Stat      `json:"shirsize,omitempty"`
	Gender                   []models.Stat      `json:"gender,omitempty"`
	GraduationYear           []models.Stat      `json:"graduationyear,omitempty"`
	IsNovice                 []models.Stat      `json:"isnovice,omitempty"`
	Status                   []models.Stat      `json:"status,omitempty"`
	Major                    []models.Stat      `json:"major,omitempty"`
	Attendees                []models.Stat      `json:"attendees,omitempty"`
}
