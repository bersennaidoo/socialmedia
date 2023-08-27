package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitemty"`
	Name      string             `json:"name,omitemty" bson:"name,omitemty"`
	Email     string             `json:"email,omitemty" bson:"email,omitemty"`
	UpdatedAt time.Time          `json:"updatedAt,omitemty" bson:"updatedAt,omitemty"`
	Password  string             `json:"password,omitemty" bson:"password,omitemty"`
}
