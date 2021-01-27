package main

import (
    "api/server"
)

func main() {
    r := server.GetRouter()
    r.Run(":8080")

}