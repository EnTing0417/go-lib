package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"encoding/json"
	"net/http"
	"github.com/EnTing0417/go-lib/model"
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
	"time"
	"golang.org/x/oauth2"

)

var (
    oauthConfig *oauth2.Config
    oauthState  = "google_login"
	config *model.Config
)

func Create(client *mongo.Client, _collection string,_model interface{}) (err error) {
	database := client.Database(DATABASE)
	collection := database.Collection(_collection)
	ctx := context.Background()

	_, err = collection.InsertOne(ctx, _model)
	if err != nil {
		return err
	}
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			return err
		}
		return err
	}
	if err := cursor.Err(); err != nil {
		return err
	}

	return nil
}

func InitGoogleAuth() {
	config = model.ReadConfig()

    oauthConfig = &oauth2.Config{
        ClientID:     config.GoogleOAuth.ClientID,
        ClientSecret: config.GoogleOAuth.ClientSecret,
        RedirectURL:  config.GoogleOAuth.RedirectUrl, 
        Scopes: []string{
            config.GoogleOAuth.Scopes[0],
            config.GoogleOAuth.Scopes[1],
        },
        Endpoint: oauth2.Endpoint{
            AuthURL:  config.GoogleOAuth.AuthURL,
            TokenURL: config.GoogleOAuth.TokenURL,
        },
    }
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
    url := oauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline)
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}


//Still in progress
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
    code := r.FormValue("code")

    token, err := oauthConfig.Exchange(context.Background(), code)
    if err != nil {
        log.Printf("Error exchanging token: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    client := oauthConfig.Client(context.Background(), token)
    
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Error getting user information: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userInfo model.User

	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		log.Printf("Error decoding user information: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	bearer_token := model.Token{
		AccessToken: token.AccessToken,
		TokenType: token.TokenType,
		RefreshToken: token.RefreshToken,
	}

	new_token, err := model.EncryptToken(bearer_token.AccessToken)

	expiration := time.Now().Add(time.Hour)

	new_bearer_token := model.AccessToken{
		ID: primitive.NewObjectID(),
		Token: new_token,
		UserID: userInfo.ID,
		Expiry: expiration,
	}

	fmt.Printf("encrypted token: %v", new_bearer_token)

	_client := Init()
	Connect(_client)

	accessToken := &model.AccessToken{
		ID : primitive.NewObjectID(),
		Token: fmt.Sprintf("%v",new_bearer_token),
		Expiry: expiration,
	}
	Create(_client,COLLECTION_ACCESS_TOKENS,accessToken)
	Disconnect(_client)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&bearer_token)
}
