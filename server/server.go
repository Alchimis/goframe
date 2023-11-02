package server

import (
	"log"
	"net/http"

	"game"

	"github.com/gorilla/websocket"
)

const (
	NETWORK           = "tcp"
	ADDRESS           = ":7070"
	READ_BUFFER_SIZE  = 1024
	WRITE_BUFFER_SIZE = 1024
)

func allowAll(req *http.Request) bool {
	return true
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  READ_BUFFER_SIZE,
	WriteBufferSize: WRITE_BUFFER_SIZE,
	CheckOrigin:     allowAll,
}

func handleMessage(message []byte) {
	log.Println(string(message))
}

/*
	 смотри если так подумать
		пользователь регистрируеться в системе и получает ключь
		по этому ключу будут рассылаться сообщения всем людям в комнате
		 вернее комната будет держат ключи всех пользователей, и когда один пользователь кидает сообщение то оно пересылаеться всем соединениям по ключам
		 оке, держать вебсокеты в памяти или в хеше
		пока есть соединение есть и пользователь. по идее ему можно сразу при подключении давать id(генериравать) а при отключении удалять
		все данные, связанные с id
*/
func HandleWebsocket(responseWriter http.ResponseWriter, request *http.Request) {
	//responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	conn, err := upgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		log.Println("Problem establishing connection: ", err)
		return
	}
	defer func() {
		conn.Close()
		log.Println("Connection are closed")
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

			handleMessage(buffer)
		case websocket.CloseMessage:
			log.Println("Conection closed")
			return
		default:
		}

	}
}

func ServerMain() {
	http.HandleFunc("/game", game.HandleResponseToGame)
	http.HandleFunc("/connect", HandleWebsocket)
	log.Println("Server started")
	err := http.ListenAndServe(ADDRESS, nil)
	if err != nil {
		log.Println("Http err", err)
	}
}

func WebsocketMain() {

}
