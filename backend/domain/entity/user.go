package entity

import (
    "time"
)

// User represents a user entity in the domain
type User struct {
    ID          uint
    FirebaseUID string
    DisplayName string
    Bio         string
    Location    string
    Website     string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
