package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateOne(client *mongo.Client, _collection string,_model interface{}) (err error) {
	database := client.Database(DATABASE)
	collection := database.Collection(_collection)
	ctx := context.Background()
	_, err = collection.InsertOne(ctx, _model)
	if err != nil {
		return err
	}
	return nil
}

func FindBy(client *mongo.Client, _collection string,criteria map[string]interface{}) ( model interface{},err error) {
	database := client.Database(DATABASE)
	collection := database.Collection(_collection)
	ctx := context.Background()

	cursor, err := collection.Find(ctx, criteria)
	if err != nil {
		return model,err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		err := cursor.Decode(&model)
		if err != nil {
			return model,err
		}
		return model,err
	}
	if err := cursor.Err(); err != nil {
		return model,err
	}
	return model,nil
}

func ListBy(client *mongo.Client, _collection string,criteria map[string]interface{},sort map[string]interface{}) ( modelList []interface{},err error) {
	database := client.Database(DATABASE)
	collection := database.Collection(_collection)
	ctx := context.Background()

	options := options.Find().SetSort(sort)

	cursor, err := collection.Find(ctx, criteria,options)
	if err != nil {
		return modelList,err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var model interface{}

		err := cursor.Decode(&model)
		if err != nil {
			return modelList,err
		}
		modelList = append(modelList,model)
	}
	if err := cursor.Err(); err != nil {
		return modelList,err
	}
	return modelList,nil
}

func UpdateBy(client *mongo.Client, _collection string, criteria map[string]interface{}, set map[string]interface{}) (model interface{},err error) {
	database := client.Database(DATABASE)
	collection := database.Collection(_collection)
	ctx := context.Background()

	_, err = collection.UpdateOne(ctx,criteria,bson.M{"$set": set})
	if err != nil {
		return model,err
	}

	cursor, err := collection.Find(ctx, criteria)
	if err != nil {
		return model,err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		err := cursor.Decode(&model)
		if err != nil {
			return model,err
		}
		return model,err
	}
	if err := cursor.Err(); err != nil {
		return model,err
	}
	return model,nil
}

//Soft delete
func DeleteBy(client *mongo.Client, _collection string, criteria map[string]interface{},set map[string]interface{}) (err error) {
	database := client.Database(DATABASE)
	collection := database.Collection(_collection)
	ctx := context.Background()

	_, err = collection.UpdateMany(ctx,criteria,set)
	if err != nil {
		return err
	}
	return nil
}

