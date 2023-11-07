package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"inter"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type SendMessagePing struct {
	ChannelId string `json:"chanel_id"`
	Message   string `json:"message"`
}

func NotifyConcreteUser(conn *UserConnection, message string) {
	conn.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

/*
я еблан
я понял это
вот в ч
нужно сообщение отправлять в канал, так правельнее
и не еби себе блять мозги слшком сильно. Ты начинал всё это как чил рофл пародию кринж
*/

type ChanelQuery interface {
	GetChanelById(string) (*inter.Channel, error)
}

func GetChanelById(cq interface{}, waw string) (*inter.Channel, error) {
	chq, ok := cq.(ChanelQuery)
	if !ok {
		return nil, errors.New("Bd isn't interface of ChanelQuery")
	}
	return chq.GetChanelById(waw)
}

type MessageQuery interface {
	InsertMessage(*inter.Message) (string, error)
}

func InsertMessage(mq interface{}, message *inter.Message) (string, error) {
	imq, ok := mq.(MessageQuery)
	if !ok {
		return "", errors.New("Bd isn't interface of MessageQuery")
	}
	return imq.InsertMessage(message)
}

func (server *Server) HandleUserTextMessage(conn *UserConnection, message *[]byte) {
	if json.Valid(*message) {
		decoder := json.NewDecoder(bytes.NewReader(*message))
		var mss *SendMessagePing
		err := decoder.Decode(mss)
		if err == nil {
			chanel, err := GetChanelById(server.Db, mss.ChannelId) //server.database.GetChanelById(mss.ChannelId)
			if err != nil {
				return
			}

			mmss := inter.CreateMessage(chanel.Id, conn.userId.String(), mss.Message)
			_, err = InsertMessage(server.Db, &mmss)
			if err != nil {
				return
			}
			chanel.NotifyAllUsers(func(id string) {
				//server.ActiveConnections[id]

				NotifyConcreteUser(server.ActiveConnections[id], mss.Message)
			})

		}
	}
}

// возможно стоит уведомить юзера когда он отключён
func (server *Server) ConnectUser(responseWriter http.ResponseWriter, request *http.Request) {
	//responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	// АПГРЕЙДЕР ВЫНЕСТИ КУДА ТО
	conn, err := server.Upgrader.Upgrade(responseWriter, request, MakeHeader())
	if err != nil {
		log.Println("Problem establishing connection: ", err)
		return
	}
	var userConn *UserConnection
	userConn, err = ConstructUserConnection(conn)
	if err != nil {
		HandleConstructError(err)
		return
	}
	server.ActiveConnections[userConn.id.String()] = userConn
	defer func() {
		conn.Close()
		log.Println("Connection are closed")
		delete(users, userConn.id)
	}()
	conn.SetCloseHandler(nil)
	log.Println("Websocket connection are established")
	buffer := make([]byte, 1024)
	var messageType int
	for {

		messageType, buffer, err = conn.ReadMessage()
		if err != nil {
			log.Println("Cant read message", err)
			return
		}
		switch messageType {
		//case websocket.PingMessage:
		//case websocket.PongMessage:
		case websocket.TextMessage:
			//case websocket.BinaryMessage:
			//conn.WriteMessage(websocket.TextMessage, buffer)

			handleMessage([]byte(userConn.id.String() + " :" + string(buffer)))
		case websocket.CloseMessage:
			log.Println("Conection closed")
			handleMessage([]byte(userConn.id.String() + " are disconected from chat"))
			return
		default:
		}

	}

}
