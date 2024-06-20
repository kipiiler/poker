package adapters

import (
	"fmt"
	service "huskyholdem/service"
	"huskyholdem/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type GameHandler struct {
	GameService *service.GameService
	channel     chan string
}

func NewGameHandler(gs *service.GameService, channel chan string) *GameHandler {
	return &GameHandler{
		GameService: gs,
		channel:     channel,
	}
}

func (h *GameHandler) CreateNewGame(c *gin.Context) {
	fmt.Println("Create new game")
	gameId := uuid.New().String()
	game, err := h.GameService.CreateNewGameWithID([]string{}, gameId)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.HandleSuccessWithMessage(c, gin.H{"game_id": game.GameID}, "success")
}

func (h *GameHandler) GameSocketByID(c *gin.Context) {
	// Check if game exists
	gameID := c.Param("room")
	_, err := h.GameService.GetGameByID(gameID)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	// start a poker game with socket if 5 player joined

	// // Upgrade the connection to a WebSocket
	// conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// defer conn.Close()

	utils.HandleSuccessWithoutData(c, "noice")

	// Handle WebSocket messages

}

func (h *GameHandler) BeNice(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	fmt.Println("Be nice")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	// Handle WebSocket messages
	for {
		// messageType, p, err := conn.ReadMessage()
		// if err != nil {
		// 	return
		// }

		// // print message in json format
		// fmt.Printf("recv: %s\n", p)

		// parse message
		type data struct {
			Name string `json:"name"`
		}

		type message struct {
			Type string `json:"type"`
			Data data   `json:"data"`
		}

		var msg message
		fmt.Println("naowidnwod")

		// parse from p to msg
		err = conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}

		h.channel <- msg.Data.Name

		fmt.Printf("recv: %s\n", msg.Data.Name)

		// err = conn.WriteMessage(messageType, p)
		// if err != nil {
		// 	return
		// }

		// send a message every 1 second
		// go func() {
		// 	for {
		// 		time.Sleep(1 * time.Second)
		// 		err = conn.WriteMessage(1, []byte("Nah bro, you be nice!")) // Convert the string to a byte slice
		// 		if err != nil {
		// 			return
		// 		}
		// 	}
		// }()
	}
}

// In-memory game hub (use to manage game room)
type GameHub struct {
	rooms   map[string]*GameRoom
	message chan string
	timer   chan string
}

func NewGameHub() *GameHub {
	return &GameHub{
		rooms: make(map[string]*GameRoom),
	}
}

func (h *GameHub) AddGameRoom(roomID string, room *GameRoom) {
	h.rooms[roomID] = room
}

func (h *GameHub) BroadCastMessage(roomID string, message string) {
	room := h.rooms[roomID]
	room.BroadCastMessageToAll(message)
}

func (h *GameHub) CreateNewGameRoom(roomID string, botIds []string, clientConns []*websocket.Conn) {
	h.rooms[roomID] = &GameRoom{
		roomID:      roomID,
		botIds:      botIds,
		clientConns: make(map[string]*websocket.Conn),
	}
}

func (h *GameHub) GetGameRoom(roomID string) *GameRoom {
	return h.rooms[roomID]
}

// In-memory game room
type GameRoom struct {
	roomID      string
	botIds      []string
	clientConns map[string]*websocket.Conn
}

func (r *GameRoom) AddBots(botId string, conn *websocket.Conn) {
	r.botIds = append(r.botIds, botId)
	r.clientConns[botId] = conn
}

func (r *GameRoom) RemoveBots(botId string) {
	for i, id := range r.botIds {
		if id == botId {
			r.botIds = append(r.botIds[:i], r.botIds[i+1:]...)
		}
	}

	delete(r.clientConns, botId)
}

func (r *GameRoom) BroadCastMessageToAll(message string) {
	for _, conn := range r.clientConns {
		conn.WriteMessage(websocket.TextMessage, []byte(message))
	}
}

func (r *GameRoom) BroadCastMessageById(message string, botId string) {
	// if botId not in room, return
	if _, ok := r.clientConns[botId]; !ok {
		r.clientConns[botId].WriteMessage(websocket.TextMessage, []byte(message))
	}
}
