package main

import (
	"fmt"
	"github.com/pwh-pwh/aiwechat-vercel/api"
	"net/http"
)

func main() {
	http.HandleFunc("/api/wx", api.Wx)
	http.HandleFunc("/api/chat", api.Chat)
	http.HandleFunc("/api/check", api.Check)
	http.HandleFunc("/api/index", api.Index)
	http.HandleFunc("/api/wx_menu", api.WxMenu)

	fmt.Println("Server is listening on port 18900...")
	err := http.ListenAndServe(":18900", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
