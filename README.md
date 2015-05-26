# Pipesock

## A websocket server you can pipe to.

![pipesockn'](http://www.spillcontainment.com/sites/default/files/pipe-light1.jpg)

Just a heads up: I'm still tweaking this a lot. If you have suggestions, problems please don't hesitate to open an issue.

### To Install:

    go get code.google.com/p/go.net/websocket
    go get github.com/minikomi/pipesock
    cd $GOPATH/src/github.com/minikomi/pipesock
    go install
    
Then, provided you have `$GOPATH/bin` in your `PATH` you can start using the server with the pipesock command:

    tail -f "some.log" | awk 'print{$2 "," $3}' | pipesock 

By default, you can then point your browser at `localhost:9193`.

### Notes:

* Other views can be added to `$GOPATH/src/github.com/minikomi/pipesock/views/`
* Accessing `/flush` will clear the buffer.
* The websocket can connect to `/ws` and receive serialized messages in the format:

    {
      Time: Timestamp for broadcast,
      Messages: [
        {
         Time: Timestamp for readline,
         Message: Message string
         }
         ....
      ]
    }

* The array of buffered messages can be read in JSON format at `/buffer.json`.

## flags:

    -d=2000: Delay between broadcasts of bundled events in ms (shorthand).
    -delay=2000: Delay between broadcasts of bundled events in ms.
    -l=false: Log HTTP requests tp STDOUT (shorthand).
    -log=false: Log HTTP requetsts to STDOUT
    -n=20: Number of previous broadcasts to keep in memory (shorthand).
    -num=20: Number of previous broadcasts to keep in memory.
    -p=9193: Port for the pipesock to sit on (shorthand).
    -port=9193: Port for the pipesock to sit on.
    -t=false: Pass output to STDOUT (shorthand).
    -through=false: Pass output to STDOUT.
    -v="default": View directory to serve (shorthand).
    -view="default": View directory to serve.

## Todo: 

* better d3 frontend
* more views
