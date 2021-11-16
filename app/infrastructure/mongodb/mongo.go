package mongodb

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MongoUrl = "mongodb+srv://dashboard_golang:SqECbvw0bEc2UsAn@cluster0.njavs.mongodb.net/dashboard?retryWrites=true&w=majority"
const Database = "dashboard"

func Connect() (client *mongo.Client) {
	clientOptions := options.Client().
		ApplyURI(MongoUrl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
