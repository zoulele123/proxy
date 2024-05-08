package pkg
import (
	"io/ioutil"
	"fmt"
	"net/http"
	"net/url"
)
func SendMetrics(serverIP, targetIP string) {
	if targetIP == "" {
		fmt.Println("Usage: go run client.go -ip <target_ip:port>")
		return
	}

	// 构建查询参数
	queryParams := url.Values{}
	queryParams.Add("ip", targetIP)
	queryString := queryParams.Encode()

	// 创建 POST 请求
	req, err := http.NewRequest("POST", "http://" + serverIP + "/metrics?" + queryString, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取服务器响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 输出服务器响应
	fmt.Println("Server response:", resp.Status)
	fmt.Println("Response body:", string(body))
}
