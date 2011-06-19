package main

import "web"
import "framework/starter"

func main() {
    web.Get("/(.*)",  starter.Get)
    web.Post("/(.*)", starter.Post)
    web.Run("0.0.0.0:8080")
}
