package websocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

	defer ws.Close()
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

	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}

		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}
