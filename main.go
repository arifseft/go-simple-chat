package main
import (
        "github.com/astaxie/beego"
        _ "github.com/arifseft/go-simple-chat/routers"
        "strconv"
        "os"
)
func main() {
        port, err := strconv.Atoi(os.Getenv("PORT"))
        if err == nil {
                beego.BConfig.Listen.HTTPPort = port
        }
        beego.Run()
}