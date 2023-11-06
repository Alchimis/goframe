package server

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"game"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	IP                = "192.168.108.194"
	PORT              = ":7070"
	NETWORK           = "tcp"
	ADDRESS           = IP + PORT //":7070"
	READ_BUFFER_SIZE  = 1024
	WRITE_BUFFER_SIZE = 1024
)

var users = make(map[uuid.UUID]*UserConnection)

func NotifyAllUsers(message []byte) {
	for _, user := range users {
		user.conn.WriteMessage(websocket.TextMessage, message)
	}
}

var channels = make(map[uuid.UUID]*Channel)

func createChannel(name string) (*Channel, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	channel := &Channel{
		UsersIds:  make([]UUIDjson, 0),
		ChannelId: id,
		Name:      name,
		CreatedAt: time.Now(),
	}
	return channel, nil
}

func createAndRegisterChannel(name string, creatorId uuid.UUID) (uuid.UUID, error) {
	channel, err := createChannel(name)
	if err != nil {
		log.Println("Eblan, kanal ne sozdalsa, vot oshibka: ", err)
		return uuid.Nil, err
	}
	channel.AddUser(creatorId)
	id := channel.ChannelId
	channels[id] = channel
	return id, nil
}

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
	NotifyAllUsers(message)
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

func MakeHeader() (header http.Header) {
	header = make(http.Header)
	header.Set("Access-Control-Allow-Origin", "*")

	return
}

type CreateChanelRequest struct {
	CreatorId   UUIDjson `json:"creator_id"`
	ChannelName string   `json:"channel_name"`
}

// POST /channel
func CreateChanel(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	switch request.Method {
	case http.MethodPost:
		var requestBody *CreateChanelRequest
		err := json.NewDecoder(request.Body).Decode(requestBody)
		if err != nil {
			log.Println("Error with create chanel decoding: ", err)
			return
		}
		createAndRegisterChannel(requestBody.ChannelName, requestBody.CreatorId.Id)

	default:
		responseWriter.Write([]byte("method " + request.Method + " not allowed"))
	}
}

func HandleWebsocket(responseWriter http.ResponseWriter, request *http.Request) {
	//responseWriter.Header().Set("Access-Control-Allow-Origin", "*")

	conn, err := upgrader.Upgrade(responseWriter, request, MakeHeader())
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
	users[userConn.id] = userConn
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
