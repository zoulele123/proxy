package pkg
import (
	"fmt"
	"net/http"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"bytes"
)

func ScriptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
			ip := r.URL.Query().Get("ip")
	fmt.Println(ip)

	if username == "" || password == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	script := string(body)

	// 建立SSH连接
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", "192.168.145.102:22", config)
	if err != nil {
		http.Error(w, "Error connecting to remote server: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// 执行远程脚本
	session, err := client.NewSession()
	if err != nil {
		http.Error(w, "Error creating SSH session: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer session.Close()

	// 将脚本内容传递给远程服务器并执行
	var stdout bytes.Buffer
	session.Stdout = &stdout
	if err := session.Run(script); err != nil {
		http.Error(w, "Error executing script on remote server: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 返回脚本执行结果
	w.Header().Set("Content-Type", "text/plain")
	w.Write(stdout.Bytes())
}
