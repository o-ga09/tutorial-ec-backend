package main

import "github.com/o-ga09/tutorial-ec-backend/app/server"

// @title ECサイトバックエンドAPI
// @version v0.1.0
// @description Goでクリーンアーキテクチャを使用して。バックエンドAPIを作成する
// @host localhost:8080
func main() {
    server, err := server.NewServer()
    if err != nil {
        panic(err)
    }
    server.Run()
}
