package http

import (
	"chapar/internals/core/domain"
	"chapar/internals/core/ports"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

// var (
// 	newline = []byte{'\n'}
// 	space   = []byte{' '}
// )

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type HTTPServer struct {
	innerBridges ports.InnerBridges
}

func NewHttpService(bridges ports.InnerBridges) HTTPServer {
	return HTTPServer{innerBridges: bridges}
}

type WebsocketConnection struct {
	id        uint
	mailman   chan domain.Message
	conn      *websocket.Conn
	terminate func(domain.User)
}

func (c *WebsocketConnection) readPump() {
	defer func() {
		c.terminate(domain.User{ID: c.id})
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, byteMessage, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message := domain.Message{}
		err = json.Unmarshal(byteMessage, &message)
		if err != nil {
			// give an invalid json signal
			log.Println(err.Error())
		}
		c.mailman <- message
	}
}

func (c *WebsocketConnection) writePump(incomingMessages chan domain.Message) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-incomingMessages:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Fatal("error : ", err) // TODO : proper error handling
				return
			}

			strMessage, _ := json.Marshal(message)
			w.Write(strMessage)

			// Add queued chat messages to the current websocket message.
			// n := len(c.send)
			// for i := 0; i < n; i++ {
			// 	w.Write(newline)
			// 	w.Write(<-c.send)
			// }

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (gs *HTTPServer) ServeWs(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	rciver := make(chan domain.Message)
	id, _ := strconv.ParseInt(c.Request().Header["Id"][0], 10, 0) // TODO : error handling
	client := &WebsocketConnection{
		id:        uint(id),
		mailman:   gs.innerBridges.Register(domain.User{ID: uint(id), Reciver: rciver}),
		conn:      conn,
		terminate: gs.innerBridges.UnRegister}

	go client.writePump(rciver)
	go client.readPump()

	return nil
}
