package mongodb

import (
	"context"
	"dashboardapi/app/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func Login(user models.User) (retval string) {
	client := Connect()
	collection := client.Database(Database).Collection("Users")

	cur, err := collection.Find(context.Background(), bson.D{
		{"username", user.Username},
		{"password", GetMD5Hash(user.Password)},
	})

	if err != nil {
		log.Fatal(err)
		retval = "Fail"
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		result := struct {
			Id string
		}{}

		err := cur.Decode(&result)

		if err != nil {
			retval = "Fail"
			log.Fatal(err)
		}

		// fmt.Println(result)
		retval = "Success"
	}

	client.Disconnect(context.Background())

	return
}
