package requests

import (
	"../models"
)

type AttendeeRegistration struct {
    Attendee     models.Attendee   `json:"attendee,omitempty"`
}
