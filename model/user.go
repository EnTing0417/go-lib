package model

import ( 
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID       `json:"_id" bson:"_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt	time.Time`json:"deleted_at" bson:"deleted_at"`
	Name       string    `json:"name" bson:"name"`
	PictureURL	   string	`json:"picture" bson:"picture"`
	Email string `json:"email"  bson:"email"`
	Status	 string `json:"status"  bson:"status"`
}