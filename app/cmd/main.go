package main

import "github.com/o-ga09/tutorial-ec-backend/app/server"

func main() {
    server, err := server.NewServer()
    if err != nil {
        panic(err)
    }
    server.Run()
}
