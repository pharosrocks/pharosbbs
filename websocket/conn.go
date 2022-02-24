package websocket

import (
	"net"
	"time"

	"github.com/gorilla/websocket"
)

type Conn struct {
	Websocket *websocket.Conn
}

func (conn *Conn) Read(b []byte) (n int, err error) {
	_, buf, err := conn.Websocket.ReadMessage()
	n = copy(b, buf)
	return n, err
}

func (conn *Conn) Write(b []byte) (n int, err error) {
	return len(b), conn.Websocket.WriteMessage(websocket.BinaryMessage, b)
}

func (conn *Conn) Close() (err error) {
	return conn.Websocket.Close()
}

func (conn *Conn) LocalAddr() (addr net.Addr) {
	return conn.Websocket.LocalAddr()
}

func (conn *Conn) RemoteAddr() (addr net.Addr) {
	return conn.Websocket.RemoteAddr()
}

func (conn *Conn) SetDeadline(t time.Time) (err error) {
	if err = conn.Websocket.SetReadDeadline(t); nil != err {
		return err
	}
	if err = conn.Websocket.SetWriteDeadline(t); nil != err {
		return err
	}
	return nil
}

func (conn *Conn) SetReadDeadline(t time.Time) (err error) {
	return conn.Websocket.SetReadDeadline(t)
}

func (conn *Conn) SetWriteDeadline(t time.Time) (err error) {
	return conn.Websocket.SetWriteDeadline(t)
}
