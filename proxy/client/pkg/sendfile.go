package pkg

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// compressFile 压缩文件
func compressFile(filePath string) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(gz, file)
	if err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func SendFile(serverIP, filePath, username, password, targetIP string) {
	// 检查参数是否完整
	if filePath == "" || username == "" || password == "" || targetIP == "" {
		fmt.Println("Usage: go run client.go -F <file_path> -username <username> -password <password> -ip <target_ip>")
		return
	}

	// 压缩文件
	compressedData, err := compressFile(filePath)
	if err != nil {
		fmt.Println("Error compressing file:", err)
		return
	}

	// 准备表单数据
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 添加压缩文件字段
	part, err := writer.CreateFormFile("file", filepath.Base(filePath) + ".gz")
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	part.Write(compressedData)

	// 添加其他字段
	writer.WriteField("username", username)
	writer.WriteField("password", password)
	writer.WriteField("ip", targetIP)

	// 关闭表单写入器
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing writer:", err)
		return
	}

	// 创建 POST 请求
	req, err := http.NewRequest("POST", "http://" + serverIP + "/upload", body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置查询参数
	q := req.URL.Query()
	q.Add("username", username)
	q.Add("password", password)
	q.Add("ip", targetIP)
	req.URL.RawQuery = q.Encode()

	// 设置请求头
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 输出服务器响应
	fmt.Println("Server response:", resp.Status)
}
