package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/Yesterday17/bili-archive-frontend/statik"
	"github.com/Yesterday17/bili-archive/components"
	"github.com/pkg/browser"
	"github.com/rakyll/statik/fs"
	"log"
	"net"
	"net/http"
)

func main() {
	handler := http.NewServeMux()
	if err := components.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	// 后端
	go components.CreateBiliArchiveServer()

	// 前端
	frontend, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	handler.Handle("/", http.FileServer(frontend))

	handler.HandleFunc("/port", func(w http.ResponseWriter, req *http.Request) {
		output, err := json.Marshal(map[string]string{"port": components.Configuration.Port})
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Write(output)
	})

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	// 打开前端网页
	log.Println(fmt.Sprintf("http://localhost:%d", listener.Addr().(*net.TCPAddr).Port))
	if err := browser.OpenURL(fmt.Sprintf("http://localhost:%d", listener.Addr().(*net.TCPAddr).Port)); err != nil {
		log.Fatal(err)
	}

	panic(http.Serve(listener, handler))
}
