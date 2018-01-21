package routers

import (
	"github.com/arifseft/go-simple-chat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.AppController{})
	
	beego.Router("/join", &controllers.AppController{}, "post:Join")

	beego.Router("/chat", &controllers.WebSocketController{})
	beego.Router("/chat/join", &controllers.WebSocketController{}, "get:Join")

}
