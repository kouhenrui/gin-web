package ws

import (
	"HelloGin/src/global"
	"HelloGin/src/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

// 定义一个 Data 结构体，用于保存用户的信息
type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	UserName string   `json:"username"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}

// 定义一个 connection 结构体，用于保存每个连接的信息
type connection struct {
	ws        *websocket.Conn // WebSocket 连接
	data      *Data           // 用户数据
	sc        chan []byte     // 发送消息的通道
	bc        chan []byte     //广播信道
	isPrivate bool
}

// 设置websocket
// CheckOrigin防止跨站点的请求伪造
var upGrader = &websocket.Upgrader{
	//设置读取写入字节大小
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		//可以添加验证信息
		return true
	},
}

// 全局消息通道，用于接收所有连接的消息
var broadcast = make(chan []byte)

type ConnectionManager struct {
	connections map[string]*connection
	lock        sync.Mutex
	broadcast   chan []byte
}

func (cm *ConnectionManager) AddConnection(conn *connection) {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	cm.connections[conn.data.Ip] = conn
}

func (cm *ConnectionManager) RemoveConnection(conn *connection) {
	delete(cm.connections, conn.data.Ip)
	close(conn.sc)
}

func (cm *ConnectionManager) Broadcast(message []byte) {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	for _, conn := range cm.connections {
		if !conn.isPrivate {
			conn.sc <- message
		}
	}
}
func (cm *ConnectionManager) StartBroadcasting() {
	for {
		select {
		case message := <-cm.broadcast:
			cm.Broadcast(message)
		}
	}
}

// 创建连接管理器实例
var connectionManager = &ConnectionManager{
	connections: make(map[string]*connection),
	broadcast:   make(chan []byte),
}

func Routers(e *gin.Engine) {
	wsGroup := e.Group("/api/ws")
	{
		wsGroup.GET("/connect", wsconnect)

	}
}

// websocket实现
func wsconnect(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, wserr := upGrader.Upgrade(c.Writer, c.Request, nil)
	if wserr != nil {
		fmt.Println("websocket连接错误")
		c.Error(errors.New(util.WEBSOCKET_CONNECT_ERROR))
		return
	}
	//isPrivate := c.Query("private") == "true"
	conn := &connection{ws: ws, data: &Data{}, sc: make(chan []byte, 1024), bc: make(chan []byte, 1024), isPrivate: false}
	// 将连接添加到连接管理器
	connectionManager.AddConnection(conn)
	//加互斥锁
	//cm.lock.Lock()
	//cm.connections[conn] = true
	//cm.lock.Unlock()
	go conn.handleConnect()

}

/*
 * @MethodName handleConnect
 * @Description 处理信道数据
 * @Author khr
 * @Date 2023/5/16 11:24
 */

func (conn *connection) handleConnect() {

	//启动异步读取和发送消息的 goroutine
	go conn.readMessages()
	go conn.writeMessages()
	go conn.broadcastMessage()
	go connectionManager.StartBroadcasting()
	//go conn.RabbitToWs()
	//go conn.WsToRabbit()
	//for {
	//	select {
	//	case message := <-conn.sc:
	//
	//		//fmt.Println(message)
	//		//if string(message) == "out" {
	//		//	conn.close()
	//		//}
	//		//if string(message) == "login" {
	//		//	_ = conn.WriteMessage(websocket.TextMessage, []byte("欢迎加入ws连接"))
	//		//}
	//		//WsCon.ws.WriteMessage(websocket.TextMessage, message)
	//		fmt.Println(string(message), "这是读取信道里面的信息")
	//
	//	case message := <-conn.fc:
	//		// 发送消息给客户端
	//		conn.sendToConsumer(message)
	//	}
	//}
}

/*
 * @MethodName readMessages
 * @Description 读取写入信道数据
 * @Author khr
 * @Date 2023/5/16 14:14
 */

func (conn *connection) readMessages() {
	for {
		_, message, err := conn.ws.ReadMessage()
		if err != nil {
			log.Println("Failed to read message from WebSocket:", err)
			break
		}
		if conn.isPrivate {
			fmt.Println("说到信息了,message:", string(message))
			//name := "ws"
			//global.Producer(name, string(message))
			conn.sc <- message
		} else {

			fmt.Println("全局信道信息")
			connectionManager.broadcast <- message
			//conn.bc <- message
		}

	}
}

/*
 * @MethodName writeMessages
 * @Description 读取写信道数据
 * @Author khr
 * @Date 2023/5/16 14:12
 */

func (conn *connection) writeMessages() {
	for message := range conn.sc {
		conn.sendToConsumer(message)
		//conn.fc <- message
		//err := conn.ws.WriteMessage(websocket.TextMessage, message)
		//if err != nil {
		//	log.Println("Failed to write message to WebSocket:", err)
		//	conn.HandleErrorMessage(err)
		//}
	}
}

/*
 * @MethodName HandleErrorMessage
 * @Description 信道数据发送错误处理
 * @Author khr
 * @Date 2023/5/16 14:14
 */

func (conn *connection) HandleErrorMessage(err error) {
	fmt.Println("error:", err)
	conn.ws.WriteMessage(websocket.TextMessage, []byte("Connection file!n"))
	conn.ws.Close()
}

func (conn *connection) sendToConsumer(out []byte) {
	err := conn.ws.WriteMessage(websocket.TextMessage, out)
	if err != nil {
		conn.HandleErrorMessage(err)
	}
}

// 广播信息
func (conn *connection) broadcastMessage() {

	for msg := range conn.bc {
		conn.sendToConsumer(msg)
	}
	//name := "ws"
	//conn.bc <- []byte(global.Consumer(name))
	//for {
	//	conn.sendToConsumer(<-conn.bc)
	//}
	//for con := range cm.connections {
	//	select {
	//	case message <- con.bc: // 发送消息到连接的通道中
	//		conn.sendToConsumer(message)
	//	default:
	//		// 如果连接的通道已满，则从连接管理器中删除该连接
	//		delete(cm.connections, con)
	//		close(conn.sc)
	//	}
	//}
}

/*
 * @MethodName WsToRabbit
 * @Description
 * @Author khr
 * @Date 2023/5/17 9:53
 */
func (conn *connection) WsToRabbit() {
	name := "ws"
	for {
		global.Producer(name, string(<-conn.sc))
	}
}

/*
 * @MethodName
 * @Description
 * @Author khr
 * @Date 2023/5/17 9:55
 */
func (conn *connection) RabbitToWs() {
	name := "ws"
	broadcast <- []byte(global.Consumer(name))
}
