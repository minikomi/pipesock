# adb logcat -> json -> websocket server

In order to facilitate making tools to work with adb output a little easier,
a small server which parses adb logcat output and then passes it to a websocket.

Still lots to do, but it's working.

Current status...

1. go get code.google.com/p/go.net/websocket
2. connect android phone
3. go run adbwebsocket.go
4. open localhost:12345 in websocket-friendly browser & watch json output
