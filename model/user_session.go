package model


import ( 
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserSession struct {
	ID          primitive.ObjectID       `json:"_id" bson:"_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt	time.Time`json:"deleted_at,omitempty" bson:"deleted_at,omitempty"`
	UserID      primitive.ObjectID     `json:"user_id" bson:"user_id"`
	Token       *Token 	`json:"token" bson:"token"`
	RefreshToken  *Token 	`json:"refresh_token" bson:"refresh_token"`
	Status	 string `json:"status"  bson:"status"`
}

