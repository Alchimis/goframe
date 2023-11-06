package main

import (
	"dbs"
	"internal"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/ostafen/clover/v2/document"
	//"server"
)

func main() {
	database := dbs.Init([]string{"users"})
	_, err := database.Run()
	if err != nil {
		log.Println("error with db running: ", err)
		return
	}

	user := &internal.User{
		Nickname:        "slava",
		Password:        "merlow",
		JoinAt:          time.Now(),
		Id:              uuid.New(),
		LiveConnections: make([]uuid.UUID, 0),
	}
	var id string
	dddd := document.NewDocumentOf(user)
	id, err = database.InsertRecord("users", dddd.AsMap())
	if err != nil {
		log.Println("Can't insert user in db, err: ", err)
		return
	}
	log.Println("id: ", id)

	var doc *document.Document
	doc, err = database.GetRecordWhere("users", map[string]interface{}{
		"nickname": "slava",
	})
	d := &struct {
		Nickname        string      `clover"nickname"`
		Password        string      `clover"password"`
		Id              uuid.UUID   `clover"id"`
		JoinAt          time.Time   `clover"join_at"`
		LiveConnections []uuid.UUID `clover"llive_conections"`
	}{}

	doc.Unmarshal(d)
	log.Println("d: ", d)
	log.Println("doc: ", doc)

	log.Println(database.GetAllUsers("users"))

	//dbs.Run()

	//server.ServerMain()
}
