# Pipesock

A websocket server you can pipe to.

To get going, create `~/.pipesock` on your home directoy, and then move the default folder inside.

Other views can also live there

    tail -f "some.log" | go run pipesock.go 
    
The server will be there on port `:9193`.

Accessing `/flush` will clear the buffer.

The websocket can connect to `/ws` and receive serialized messages in the format:

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

The array of buffered messages can be read in JSON format at `/buffer.json`.

![pipesockn'](http://www.westernsafety.com/ultratech2008/UltratechStormpg18-PipeSock.jpg)

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
