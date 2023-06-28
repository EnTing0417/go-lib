package model

import ( 
	"go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

type Token struct {
    AccessToken  string    `json:"access_token" bson:"access_token"`
    CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt	time.Time`json:"deleted_at" bson:"deleted_at"`
    TokenType    string    `json:"token_type,omitempty" bson:"token_type,omitempty"`
    RefreshToken string    `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
    Expiry       string `json:"expire_in"  bson:"expire_in"`
}

type AccessToken struct {
	ID          primitive.ObjectID       `json:"_id" bson:"_id"`
    CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
	DeletedAt	time.Time`json:"deleted_at" bson:"deleted_at"`
    Token     string                  `json:"token" bson:"token"`
    UserID      primitive.ObjectID       `json:"user_id" bson:"user_id"`
    Expiry       time.Time `json:"expiry"  bson:"expiry"`
}

//TO-DO : ValidateToken