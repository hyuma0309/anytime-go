package main

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "api/server"
)

func main() {
    r := server.GetRouter()
    r.Run(":8080")

}

type User struct {
    gorm.Model
    NickName string `json:"nickName"`
}



func DBConnect() *gorm.DB {
    DBMS := "mysql"
    USER := "root"
    PASS := "password"
    PROTOCOL := "tcp(mysql:3306)"
    DBNAME := "sample"
    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
        panic(err.Error())
    }
    return db
}