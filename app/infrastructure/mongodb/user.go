package mongodb

import (
	"context"
	"dashboardapi/app/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) (retval string) {
	client := Connect()
	collection := client.Database(Database).Collection("Users")

	user.Password = GetMD5Hash(user.Password)

	result, insertErr := collection.InsertOne(context.Background(), user)

	if insertErr != nil {
		retval = insertErr.Error()
	} else {
		retval = result.InsertedID.(primitive.ObjectID).String()
	}

	client.Disconnect(context.Background())

	return
}
