package model

import ( 
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID       `json:"_id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt	time.Time`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	Completed   bool      `json:"completed" bson:"completed"`
	UserID		primitive.ObjectID	  `json:"user_id" bson:"user_id"`
}

