package root

import (
	/*"encoding/json"
	"net/http"*/

	"app/controllers"
	// "github.com/astaxie/beego"
	// "github.com/gorilla/websocket"
	//"app/models"
)

type ChatroomController struct {
	controllers.BaseController
}

func (this *ChatroomController) Get() {
	this.TplNames = "root/chatroom.html"
}
