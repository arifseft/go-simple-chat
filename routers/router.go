package routers

import (
	"github.com/astaxie/beego"
	"github.com/arifseft/go-simple-chat/controllers"
)

func init() {
	beego.Router("/", &controllers.AppController{})
	
	beego.Router("/join", &controllers.AppController{}, "post:Join")

	beego.Router("/chat", &controllers.WebSocketController{})
	beego.Router("/chat/join", &controllers.WebSocketController{}, "get:Join")

}
