package main

import (
	"dbs"
	internal "inter"
	"log"
	"reflect"
	"time"

	"server"

	"github.com/google/uuid"
	"github.com/ostafen/clover/v2/document"
	//"server"
)

func buh() {
	database := dbs.Init( /*[]string{"users", ""}*/ dbs.GetColections())
	_, err := database.Run()
	if err != nil {
		log.Println("error with db running: ", err)
		return
	}

	user := &internal.User{
		Nickname:        "slava",
		Password:        "merlow",
		JoinAt:          time.Now(),
		Id:              uuid.New().String(),
		LiveConnections: make([]string, 0),
	}
	var id string
	//json.Marshal(user)
	dddd := document.NewDocumentOf(user)

	//log.Println("dddd: ", dddd)
	//fff := &internal.User{}
	//dddd.Unmarshal(fff)
	//log.Println("fff: ", fff)
	id, err = database.Db.InsertOne("users", dddd) //database.InsertFocusRecord("users", dddd)
	if err != nil {
		log.Println("Can't insert user in db, err: ", err)
		return
	}
	log.Println("id: ", id)
	log.Println(database.Db.FindById("users", id))

	/*var doc *document.Document
	doc, err = database.GetRecordWhere("users", map[string]interface{}{
		"nickname": "slava",
	})*/

	/*log.Println("doc: ", doc)
	ids := doc.Get("id").([]byte)
	vv, _ := uuid.FromBytes(ids)
	log.Println("id: ", vv)*/
	/*d := &struct {
		Nickname string `clover:"nickname"`
		Password string `clover:"password"`
		//Id              uuid.UUID   `clover"id"`
		JoinAt          time.Time   `clover:"join_at"`
		LiveConnections []uuid.UUID `clover:"llive_conections"`
	}{}

	err = doc.Unmarshal(d)
	if err != nil {
		log.Println("Error with doc unmarshal: ", err)
	}
	log.Println("d: ", d)

	jj := &struct {
		Name string    `clover:"name"`
		Id   uuid.UUID `clover:"id,string"`
	}{
		Name: "jobs",
		Id:   uuid.New(),
	}

	docum := document.NewDocumentOf(jj)

	log.Println(docum)

	bbb := &struct {
		Name string    `clover:"name"`
		Id   uuid.UUID `clover:"id,string"`
	}{}

	docum.Unmarshal(bbb)*/
	/*
		log.Println(bbb)
		usdd, err := database.GetAllUsers("users")
		log.Println("get all users error: ", err)
		log.Print("Users ")
		for _, v := range usdd {
			log.Print(*v, " ")
		}
		log.Println()*/
	/*data, err := database.GetAllFocusUser("users")
	for _, val := range data.([]*document.Document) {

		g := &struct {
			Nickname        string      `clover:"nickname"`
			Password        string      `clover:"password"`
			Id              string      `clover:"id"`
			JoinAt          time.Time   `clover:"join_at"`
			LiveConnections []uuid.UUID `clover:"llive_conections"`
		}{}

		val.Unmarshal(g)
		log.Println(g)
		log.Println(uuid.Parse(g.Id))
	}*/

	//database.DeleteAllWhere("users", map[string]interface{}{
	//	"Nickname": "bebre",
	//})

	/*data, err := database.GetAllFocusUserWhere("users", map[string]interface{}{
		"Nickname": "slava",
	})*/
	conditions := map[string]interface{}{
		"Nickname": "slava",
	}
	data, err := database.GetAllInterfaceUserWhere("users", func(doc *document.Document) bool {
		for key := range conditions {
			if !reflect.DeepEqual(doc.Get(key), conditions[key]) {
				return false
			}
		}
		return true
	})
	for _, val := range data.([]*document.Document) {

		g := &struct {
			Nickname        string      `clover:"nickname"`
			Password        string      `clover:"password"`
			Id              string      `clover:"id"`
			JoinAt          time.Time   `clover:"join_at"`
			LiveConnections []uuid.UUID `clover:"llive_conections"`
		}{}

		val.Unmarshal(g)
		log.Println(g)
		log.Println(uuid.Parse(g.Id))
	}

	//dbs.Run()

	//server.ServerMain()
}

func main() {
	server.NewServerMain()
	/*
		запросы к бд.
			получение всех каналов
			получение всех пользователей из канала
			создание пользователя
			добавление пользователя в канал
			удаление пользователя из канала
			удаление пользователя нахуй
			удаление канала
			аутентификация
	*/
	//buh()
	/*jj := &struct {
		Name string    `clover:"name"`
		Id   uuid.UUID `clover:"id,string"`
	}{
		Name: "jobs",
		Id:   uuid.New(),
	}

	docum := document.NewDocumentOf(jj)

	log.Println("docum: ", docum)

	bbb := &struct {
		Name string    `clover:"name"`
		Id   uuid.UUID `clover:"id,string"`
	}{}

	docum.Unmarshal(bbb)

	log.Println("bbb: ", bbb)*/
	//server.ServerMain()

}
