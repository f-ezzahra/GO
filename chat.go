package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	connHost = "127.0.0.1"
	connPort = "135"
	connType = "tcp"
)

type Trainer struct {
	Content   string `bson:"content"`   // Contenu
	StartTime int64  `bson:"startTime"` // Temps de création
	EndTime   int64  `bson:"endTime"`   // Date d'expiration
	Read      uint   `bson:"read"`      // Lire
}

func creatconversation(database string, id string, content string, read uint, expire int64) (err error) {

	collection := conf.MongoDBClient.Database(database).Collection(id)
	comment := ws.Trainer{

		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	_, err = collection.InsertOne(context.TODO(), comment)
	return
}
func ii() {
	fmt.Println("Connecting to " + connType + " server " + connHost + ":" + connPort)
	conn, err := net.Dial(connType, connHost+":"+connPort)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")
		reader.ReadString('\n')

		message, _ := bufio.NewReader(conn).ReadString('\n')

		log.Print("Server relay:", message)
	}
}
