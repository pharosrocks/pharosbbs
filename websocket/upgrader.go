package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Upgrader struct {
	upgrader websocket.Upgrader

	Error func(w http.ResponseWriter, r *http.Request, status int, reason error)

	CheckOrigin func(r *http.Request) bool
}

func (u *Upgrader) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (conn *Conn, err error) {
	u.upgrader.CheckOrigin = u.CheckOrigin
	u.upgrader.Error = u.Error
	websocket, err := u.upgrader.Upgrade(w, r, responseHeader)

	conn = &Conn{
		Websocket: websocket,
	}

	return conn, err
}
