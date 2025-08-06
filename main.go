package main

import (
	"github.com/gin-gonic/gin"
	"github.com/andr1ww/odin"
	"log"
	"fmt"
)

const (
	port = ":3551"

)



type User struct {
    odin.Bucket `bucket:"users" database:"main"`
    Name        string `json:"name"`
    Email       string `json:"email"`
}

func main() {
    if err := odin.ConnectDefault("./odin.db"); err != nil {
        log.Fatal(err)
    }
    defer odin.CloseAll()

    user := &User{
        Bucket: odin.Bucket{ID: "Key"},
        Name:   "Luna",
        Email:  "Luna@Forge.com",
    }

    if err := odin.Create(user); err != nil {
        log.Fatal(err)
    }
}

func Routes(r *gin.Engine ) {
	fmt.Println("Forge Backend created by Luna")
}