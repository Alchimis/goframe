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
	chq, ok := cq.(*ChanelQuery)
	if !ok {
		return nil, errors.New("Bd isn't interface of ChanelQuery")
	}
	return (*chq).GetChanelById(waw)
}

type MessageQuery interface {
	InsertMessage(*inter.Message) (string, error)
}

func InsertMessage(mq interface{}, message *inter.Message) (string, error) {
	imq, ok := mq.(*MessageQuery)
	if !ok {
		return "", errors.New("Bd isn't interface of MessageQuery")
	}
	return (*imq).InsertMessage(message)
}

func (server *Server) HandleUserTextMessage(conn *UserConnection, message *[]byte) {
	if json.Valid(*message) {
		decoder := json.NewDecoder(bytes.NewReader(*message))
		var mss *SendMessagePing
		err := decoder.Decode(mss)
		if err != nil {
			log.Println("Errorr with decoding")
			return
		}

		chanel, err := GetChanelById(server.Db, mss.ChannelId) //server.database.GetChanelById(mss.ChannelId)
		if err != nil {
			log.Println("Chanel not finded, err: ", err)
			return
		}

		mmss := inter.CreateMessage(chanel.Id, conn.userId.String(), mss.Message)
		_, err = InsertMessage(server.Db, &mmss)
		if err != nil {
			log.Println("Can't insert message: ", err)
			return
		}
		chanel.NotifyAllUsers(func(id string) {
			NotifyConcreteUser(server.ActiveConnections[id], mss.Message)
		})

	}
}

const PLEASE_DONT_SEND_BINARY_STRING = "Please don't send binar"

var PLEASE_DONT_SEND_BINARY_SLICE = []byte(PLEASE_DONT_SEND_BINARY_STRING)

func (server *Server) HandleBinaryMessage(conn *UserConnection, message *[]byte) {
	conn.conn.WriteMessage(websocket.BinaryMessage, PLEASE_DONT_SEND_BINARY_SLICE)
}

func (server *Server) HandleRegistration(responseWriter http.ResponseWriter, request *http.Request) {
	//

	reader, err := request.GetBody()
	if err != nil {
		log.Println("Error with getting request body: ", err)
		// TODO: отправлять пользоателю ошибку(внутренняя ошибка сервера)
		return
	}
	buff := make([]byte, 1024)
	n, err := reader.Read(buff)
	if err != nil {
		log.Println("Error with getting request body: ", err)
		// TODO: отправлять пользоателю ошибку(внутренняя ошибка сервера)
		return
	}
	// TODO: поправить эту сепердибильную тему
	slc := buff[:n]
	var registrationRequset *RegistrationRequest
	err = json.Unmarshal(slc, registrationRequset)
	if err != nil {
		log.Println("Error with registration request unmarshal: ", err)
		// TODO: вот тут я хз. надо посмотреть какие там могут быть ошибки, потому что варианта 2: либу у меня на сервере что то произошло, либо запрос некоректный. а при некоректном запросе надо отправлять пользователю с чем именно проблема
		return
	}
	// TODO: добавить добавление нового пользователя в бд
	// TODO: сделать приколы с токенами
	// TODO: отправить токен в качестве ответа

}

func (server *Server) HandleLogin(responseWriter http.ResponseWriter, request *http.Request) {
	//
}

// возможно стоит уведомить юзера когда он отключён
// БАЗА: АУТЕНТИФИКАЦИЯ ПРОХОДИТ КАК ОБЫЧНО ЧЕРЕЗ HTTPS. В ОТВЕТЕ ОТ СЕРВЕРА ДОЛЖЕН БЫТЬ ПРОПУСК. ЭТОТ ПРОПУСК СТОИТ ИСПОЛЬЗОВАТЬ КАК ЧАСТЬ INITIAL HANDSHAKE
// И ЕЩЁ: ИСПОЛЬЗОВАТЬ WWS, ПрОВЕРЯТЬ ДАННЫЕ КАК НА ВХОДЕ ТАК И НА ВЫХОДЕ
// https://devcenter.heroku.com/articles/websocket-security
func (server *Server) ConnectUser(responseWriter http.ResponseWriter, request *http.Request) {
	//responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	// АПГРЕЙДЕР ВЫНЕСТИ КУДА ТО

	conn, err := server.Upgrader.Upgrade(responseWriter, request, MakeHeader())
	if err != nil {
		log.Println("Problem establishing connection: ", err)
		return
	}

	// TODO: добавить проверку токена регистрации
	// TODO: достать токен из бд
	var userConn *UserConnection
	userConn, err = /*ConstructUserConnectionWithUserId(conn)*/ ConstructUserConnection(conn)
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
		case websocket.BinaryMessage:
			server.HandleBinaryMessage(userConn, &buffer)
		case websocket.TextMessage:
			server.HandleUserTextMessage(userConn, &buffer)
		case websocket.CloseMessage:
			log.Println("Conection closed")
			handleMessage([]byte(userConn.id.String() + " are disconected from chat"))
			return
		default:
		}

	}

}
