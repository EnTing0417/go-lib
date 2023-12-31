package model

import ( 
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID       `json:"_id" bson:"_id"`
	Password 	string	`json:"password" bson:"password"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt	time.Time`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	Username       string    `json:"username" bson:"username"`
	PictureURL	   string	`json:"picture,omitempty" bson:"picture,omitempty"`
	Email string `json:"email,omitempty"  bson:"email,omitempty"`
	Status	 string `json:"status,omitempty"  bson:"status,omitempty"`
}