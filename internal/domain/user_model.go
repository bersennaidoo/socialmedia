package domain

import "time"

type User struct {
	Name      string    `json:"name" bson:"name"`
	Email     string    `json:"email" bson:"email"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	Password  string    `json:"password" bson:"password"`
}
