package mongodb

import (
	"context"
	"dashboardapi/app/models"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(user models.User) (response models.ResponseModel) {
	response.ResCode = -1
	response.ResMessage = "Error"
	response.ResData = "User does not exist"

	client := Connect()
	collection := client.Database(Database).Collection("Users")

	filter := bson.D{
		primitive.E{Key: "username", Value: user.Username},
		primitive.E{Key: "password", Value: GetMD5Hash(user.Password)},
	}

	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
		response.ResCode = -1
		response.ResMessage = "Error"
		response.ResData = err.Error()
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		result := struct {
			Id string
		}{}

		err := cur.Decode(&result)

		if err != nil {
			response.ResCode = -1
			response.ResMessage = "Error"
			response.ResData = err.Error()
			log.Fatal(err)
		}

		user, err := createToken(result.Id)

		if err != nil {
			response.ResCode = -1
			response.ResMessage = "Error"
			response.ResData = err.Error()
			log.Fatal(err)
		}

		// fmt.Println(result)
		response.ResCode = 0
		response.ResMessage = "Success"
		response.ResData = user
	}

	client.Disconnect(context.Background())

	return
}

func createToken(userid string) (models.UserTokens, error) {
	userToken := models.UserTokens{UserID: userid, Token: ""}

	os.Setenv("ACCESS_SECRET", "ercag") //this should be in an env file
	exp := time.Now().Add(time.Minute * 15).Unix()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = exp
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return userToken, err
	}

	ct := time.Now()

	userToken.Token = token
	userToken.Exp = exp
	userToken.CreatedDate = fmt.Sprintf("%d/%d/%d %d:%d:%d", ct.Day(), ct.Month(), ct.Year(), ct.Hour(), ct.Minute(), ct.Second())

	client := Connect()
	collection := client.Database(Database).Collection("UserTokens")

	collection.InsertOne(context.Background(), userToken)

	return userToken, nil
}

func ValidateToken(token string) (valid bool) {
	valid = false

	token = strings.Replace(token, "Bearer", "", 1)
	token = strings.TrimSpace(token)

	client := Connect()
	collection := client.Database(Database).Collection("UserTokens")

	filter := bson.D{
		primitive.E{Key: "token", Value: token},
	}

	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
		valid = false
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		result := struct {
			userid string
		}{}

		err := cur.Decode(&result)

		if err != nil {
			valid = false
			log.Fatal(err)
		}

		// fmt.Println(result)
		valid = true
	}
	client.Disconnect(context.Background())

	return valid
}
