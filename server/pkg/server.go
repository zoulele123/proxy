package pkg
import (
        "fmt"
        "io"
        "net/http"
        "os"
        "golang.org/x/crypto/ssh"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码参数
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	destIP := r.URL.Query().Get("ip")

	// 解析请求中的文件
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 创建目标文件
	dst, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 将上传的文件内容拷贝到目标文件中
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully!")

	// 如果用户名和密码均不为空，则执行 SCP 命令传输文件
	if username != "" && password != "" {
		clientConfig := &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥验证
		}

		// 连接SSH服务器
		conn, err := ssh.Dial("tcp", destIP+":22", clientConfig)
		if err != nil {
			fmt.Println("Error connecting to SSH server:", err)
			return
		}
		defer conn.Close()

		// 创建新的SSH会话
		session, err := conn.NewSession()
		if err != nil {
			fmt.Println("Error creating SSH session:", err)
			return
		}
		defer session.Close()

		// 使用SCP将文件上传到目标服务器
		go func() {
			scpWriter, _ := session.StdinPipe()
			defer scpWriter.Close()
			file, _ := os.Open(header.Filename)
			defer file.Close()
			fileInfo, _ := file.Stat()
			fmt.Fprintln(scpWriter, "C0644", fileInfo.Size(), header.Filename)
			io.Copy(scpWriter, file)
			fmt.Fprint(scpWriter, "\x00")
		}()

		// 执行SCP命令
		if err := session.Run("/usr/bin/scp -t /tmp"); err != nil {
			fmt.Println("Error transferring file via SCP:", err)
			return
		}

		fmt.Println("File transferred successfully to", destIP)
	}
}
