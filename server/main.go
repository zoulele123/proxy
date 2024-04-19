package main

import (
	"dev/pkg"
	"net/http"
	"fmt"
)
func main() {
	http.HandleFunc("/upload", pkg.UploadHandler)
	http.HandleFunc("/metrics", pkg.HandleMetricsForwarding)
	http.HandleFunc("/script", pkg.ScriptHandler)

	// 启动服务器
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
