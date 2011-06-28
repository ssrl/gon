package main

import "web"
import "framework/starter"

func main() {
    starter.Start()
    web.Config.StaticDir = "web-app/"
    web.Post("/(.*)", starter.Get)
    web.Get ("/(.*)", starter.Get)
    web.Run("0.0.0.0:8080")
}
