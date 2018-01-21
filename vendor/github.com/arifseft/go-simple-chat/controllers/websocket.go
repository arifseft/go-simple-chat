package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/arifseft/go-simple-chat/models"
)

type WebSocketController struct {
	baseController
}

func (this *WebSocketController) Get() {
	uname := this.GetString("uname")
	group := this.GetString("group")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	this.TplName = "websocket.html"
	this.Data["UserName"] = uname
	this.Data["Group"] = group
	this.Data["IsWebSocket"] = true
}

func (this *WebSocketController) Join() {
	uname := this.GetString("uname")
	group := this.GetString("group")
	if len(uname) == 0 {
		this.Redirect("/", 302)
		return
	}

	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	Join(uname, ws, group)
	defer Leave(uname)

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		publish <- newEvent(models.EVENT_MESSAGE, uname, string(p), group)
	}
}

func broadcastWebSocket(event models.Event) {
	data, err := json.Marshal(event)
	if err != nil {
		beego.Error("Fail to marshal event:", err)
		return
	}

	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		ws := sub.Value.(Subscriber).Conn
		if ws != nil {
			if ws.WriteMessage(websocket.TextMessage, data) != nil {
				unsubscribe <- sub.Value.(Subscriber).Name
			}
		}
	}
}
