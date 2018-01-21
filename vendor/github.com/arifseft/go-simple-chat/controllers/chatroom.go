package controllers

import (
	"container/list"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/arifseft/go-simple-chat/models"
)

type Subscription struct {
	Archive []models.Event
	New     <-chan models.Event
}

func newEvent(ep models.EventType, user, msg string, group string) models.Event {
	return models.Event{ep, user, int(time.Now().Unix()), msg, group}
}

func Join(user string, ws *websocket.Conn, group string) {
	subscribe <- Subscriber{Name: user, Conn: ws, Group: group}
}

func Leave(user string) {
	unsubscribe <- user
}

type Subscriber struct {
	Name string
	Conn *websocket.Conn
	Group string
}

var (
	subscribe = make(chan Subscriber, 10)
	unsubscribe = make(chan string, 10)
	publish = make(chan models.Event, 10)
	subscribers = list.New()
)

func chatroom() {
	for {
		select {
		case sub := <-subscribe:
			if !isUserExist(subscribers, sub.Name) {
				subscribers.PushBack(sub)
				publish <- newEvent(models.EVENT_JOIN, sub.Name, "", sub.Group)
				beego.Info("New user:", sub.Name, ";WebSocket:", sub.Conn != nil, ";Group:", sub.Group)
			} else {
				beego.Info("Old user:", sub.Name, ";WebSocket:", sub.Conn != nil, ";Group:", sub.Group)
			}
		case event := <-publish:
			
			broadcastWebSocket(event)
			models.NewArchive(event)

			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content, ";Group:", event.Group)
			}
		case unsub := <-unsubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					subscribers.Remove(sub)
					
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub)
					}
					publish <- newEvent(models.EVENT_LEAVE, unsub, "", sub.Value.(Subscriber).Group)
					break
				}
			}
		}
	}
}

func init() {
	go chatroom()
}

func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}
