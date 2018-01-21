package main
import (
        "github.com/astaxie/beego"
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