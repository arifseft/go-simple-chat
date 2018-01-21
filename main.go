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
        //host, err := strconv.Atoi(os.Getenv("HOST"))
        //if err == nil {
                beego.BConfig.Listen.HTTPAddr = os.Getenv("HOST")
        //}
        beego.Run()
}