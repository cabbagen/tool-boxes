package proxy

import (
	"fmt"
	"log"
	"net/http"
)

var port = ":8089"

func RunHttpProxy() {
	fmt.Printf("http 代理服务开始，监听端口为 %s...\n", port)

	proxy := NewCproxy()

	http.Handle("/", proxy)

	log.Fatal(http.ListenAndServe(port, nil))
}
