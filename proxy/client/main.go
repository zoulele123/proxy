package main
import (
	"fmt"
	"proxy/client/pkg"
	"flag"
)

func main() {
	var (
		username string
		password string
                serverIP string
	)
	// 解析命令行参数
	repath := flag.String("m", "", "metrics、script、upload")
	filePath := flag.String("s", "", "src file path")
//	username := flag.String("u", "", "Receiving server username")
//	password := flag.String("p", "", "Receiving server password")
	targetIP := flag.String("ip", "", "Receiving server target IP or targetIP:port")
	flag.Parse()
//主机账号密码
username = "root"
password = "1"
serverIP = "127.0.0.1:8080"

	// 发送请求
	switch *repath {
	case "upload":
		pkg.SendFile(serverIP, *filePath, username, password, *targetIP)
	case "script":
		pkg.SendScript(serverIP, *filePath, username, password, *targetIP)
	case "metrics":
		pkg.SendMetrics(serverIP, *targetIP)
            default:
            fmt.Println("参数请求未添加: -R")
    }
}
