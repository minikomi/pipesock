package main

import (
	"bufio"
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func IndexServer(w http.ResponseWriter, req *http.Request) {
	log.Println("Path:", req.URL.Path)
	if req.URL.Path == "/" {
		http.ServeFile(w, req, "index.html")
	} else {
		http.ServeFile(w, req, "."+req.URL.Path)
	}
}

type Hub struct {
	Connections map[*Socket]bool
	Pipe        chan []byte
}

func (h *Hub) Broadcast() {
	for {
		var x string
		x = string(<-h.Pipe)
		for s, _ := range h.Connections {
			err := websocket.Message.Send(s.Ws, x)
			if err != nil {
				log.Println(err)
				s.Ws.Close()
				delete(h.Connections, s)
			}
		}
	}
}

func (s *Socket) ReceiveMessage() {
	websocket.Message.Send(s.Ws, "Welcome")
	for {
		var x []byte
		err := websocket.Message.Receive(s.Ws, &x)
		if err != nil {
			break
		}
		h.Pipe <- x
	}
	s.Ws.Close()
}

var h Hub

type Socket struct {
	Ws *websocket.Conn
}

func wsServer(ws *websocket.Conn) {
	s := &Socket{ws}
	h.Connections[s] = true
	s.ReceiveMessage()
}

func main() {

	h.Connections = make(map[*Socket]bool)
	h.Pipe = make(chan []byte, 1)
	go h.Broadcast()

	go (func() {
		// Reader for stdin
		r := bufio.NewReader(os.Stdin)
		for {
			//Read
			str, err := r.ReadString('\n')
			if err != nil {
				log.Println("Read Error:", err)
				continue
			}
			if len(str) == 0 {
				continue
			}

			//Marshal
			msg, err := json.Marshal(map[string]string{"message": str})
			if err != nil {
				log.Println("JSON Error:", err)
				continue
			}

			//Broadcast
			h.Pipe <- (msg)
		}
	})()

	http.Handle("/ws", websocket.Handler(wsServer))
	http.HandleFunc("/", IndexServer)

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
