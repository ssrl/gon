package main

import "web"
import "framework/starter"

func main() {
    web.Config.StaticDir = "web-app/"
    web.Get("/(.*)", starter.Get)
    web.Run("0.0.0.0:8080")
}
