package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID  `bson:"_id,omitempty"`
	Phone      string `bson:"phone"`
}
type SmsVerify struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Phone      string    `bson:"phone"`
	Code       string    `bson:"code"`
	CreatedAt  time.Time 	`bson:"created_at"`
}
