package pkg
import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"fmt"
)
func HandleMetricsForwarding(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling metrics forwarding request...")

	// 获取查询参数中的目标 IP 地址
	queryIP := r.URL.Query().Get("ip")
	if queryIP == "" {
		http.Error(w, "Missing target IP in query parameters", http.StatusBadRequest)
		return
	}

	// 创建目标URL
	targetURL, err := url.Parse("http://" + queryIP)
	if err != nil {
		http.Error(w, "Invalid target URL", http.StatusInternalServerError)
		return
	}

	// 创建代理
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// 代理请求
	proxy.ServeHTTP(w, r)
}
