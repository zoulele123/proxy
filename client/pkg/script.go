package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func SendScript(filePath, username, password, targetIP string) {
	// 检查必要参数
	if username == "" || password == "" || filePath == "" || targetIP == "" {
		log.Fatal("Username, password, script file path, and IP are required")
	}

	// 读取脚本文件内容
	script, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading script file: %v", err)
	}

	// 代理服务器地址
	proxyURL := "http://localhost:8080/script"

	// 发送 POST 请求到代理服务器
	resp, err := http.Post(proxyURL+"?username="+username+"&password="+password+"&ip="+targetIP, "text/plain", bytes.NewReader(script))
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// 打印响应结果
	fmt.Println("Response:")
	fmt.Println(string(respBody))
}
