package logic

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const FieldID = "_id"
const FieldUserID = "user_id"
const FieldMoney = "money"

type MongoClient struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoClient(serverDB string, dbName string) (*MongoClient, error) {
	opts := options.Client().ApplyURI(serverDB)
	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	return &MongoClient{client: client, db: db}, nil
}

func (client *MongoClient) Disconnect() error {
	return client.client.Disconnect(context.TODO())
}

func (client *MongoClient) FindUser(userID string, ctx context.Context, collection string) (*User, error) {
	var filter bson.D

	if len(userID) == 0 {
		filter = bson.D{{FieldUserID, userID}}
	}

	var user *User
	err := client.db.Collection(collection).FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (client *MongoClient) CreateUser(userID string, ctx context.Context, collection string) error {

	_ , err := client.FindUser(userID, ctx, collection)
	if err != mongo.ErrNoDocuments || err != nil{
		return err
	}

	var user *User

	user.Id = userID

	_, err = client.db.Collection(collection).InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (client *MongoClient) UpdateUserMoney(userID string, money int64, ctx context.Context, collection string) error {
	updateTo := bson.D{
		{"$set", bson.D{{FieldMoney, money},}},
	}

	filter := bson.D{{FieldUserID, userID}}


	_, err := client.db.Collection(collection).UpdateOne(ctx, filter, updateTo)
	if err != nil || err == mongo.ErrNilDocument{
		return errors.New("update money err - "+err.Error())
	}

	return nil
}
