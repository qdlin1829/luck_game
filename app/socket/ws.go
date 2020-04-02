package socket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

//本核心在于形成userid和Node的映射关系
type Node struct {
	Conn *websocket.Conn
	//并行转串行,
	DataQueue chan []byte
}

//映射关系表
var clientMap map[int64]*Node = make(map[int64]*Node, 0)


//读写锁
var rwlocker sync.RWMutex

func Ws(c *gin.Context){
	ws,err := upGrader.Upgrade(c.Writer, c.Request, nil )
	if err != nil {
		return
	}

	//defer ws.Close()
	id := c.Query("id")
	userId, _ := strconv.ParseInt(id, 10, 64)


	//todo 获得conn
	node := &Node{
		Conn:      ws,
		DataQueue: make(chan []byte, 50),
	}

	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()

	go readMessage(node)

	go writeMessage(node)
	//for {
	//	//读取ws中的数据
	//	mt, message, err := ws.ReadMessage()
	//	if err != nil {
	//		break
	//	}
	//	if string(message) == "ping" {
	//		message = []byte("pong")
	//	}
	//
	//	//写入ws数据
	//	fmt.Println(string(message))
	//	err = ws.WriteMessage(mt, message)
	//	if err != nil {
	//		break
	//	}
	//}
}

func readMessage(node *Node){

	for {
		_, message, err := node.Conn.ReadMessage()
		fmt.Println(message)
		if err != nil {
			break
		}

		node.DataQueue <- message
		//node.Conn.WriteMessage(socket.TextMessage, message)

	}
}

type Message struct {
	UsId int64 `json:"usid"`
	DsId int64 `json:"dsid"`
	Type string `json:"type"`
	Content interface{} `json:"content"`
}

func writeMessage(node *Node){
	for {
		select {
		case data := <-node.DataQueue:
			message := &Message{}
			json.Unmarshal(data, message)
			fmt.Println(message)
			msg,_ :=json.Marshal(message)
			fmt.Println(msg)
			//err := node.Conn.WriteMessage(socket.TextMessage, data)
			err := clientMap[message.DsId].Conn.WriteMessage(websocket.TextMessage, msg)
			//clientMap[2].Conn.WriteMessage(socket.TextMessage, []byte("ok"))
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}

}
