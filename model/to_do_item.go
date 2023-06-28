package model

import ( 
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoItem struct {
	ID          primitive.ObjectID       `json:"_id" bson:"_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt	time.Time`json:"deleted_at" bson:"deleted_at"`
	Completed   bool      `json:"completed" bson:"completed"`
	UserID		string	  `json:"user_id" bson:"user_id"`
}

