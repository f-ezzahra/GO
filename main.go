package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://root:example@0.0.0.0:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("conversation").Collection("chat2")
}

type Message struct {
	ID   primitive.ObjectID `bson:"_id"`
	Text string             `bson:"text"`
}

func creatConversation(task *Message) error {
	_, err := collection.InsertOne(ctx, task)
	return err
}

func sendMessage() {
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("Text to send: ")

		a, _ := reader.ReadString('\n')
		switch {
		case a == "bye":
			fmt.Println("Good bye")
			break
		case a == "quit":
			fmt.Println("Good bye")
			break

		default:
			log.Print("you write :", a)
			str := a
			task := &Message{
				ID:   primitive.NewObjectID(),
				Text: str,
			}
			creatConversation(task)

		}

	}

}
func main() {

	sendMessage()

}
