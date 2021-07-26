# Tutorial Socket.IO

## Go
### main.go
    // Initialize functions for push
	server := socketio.NewServer(nil)

    // Setup url chat server
    server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		so.Join("chat")
		uadmin.Trail(uadmin.DEBUG, "socket onConnect remoteAddr: %v", so.RemoteAddr())
		return nil
	})

    // Setup Event
    server.OnEvent("/", "chat message", func(so socketio.Conn, msg interface{}) {
		fmt.Println("emit:", msg)

        // Send back to the client accessing the room event
		server.BroadcastToRoom("", "chat", "chat message", msg)
	})

    // Setup Event
    //Get from and event then send back to other event
    server.OnEvent("/", "updateMonitoring", func(so socketio.Conn, msg interface{}) {

        // Send Something from the backend 
		so.Emit("NotifyUpdate", msg)

        //redirect to other event
		server.BroadcastToRoom("", "chat", "NotifyUpdate", msg)
	})




### HTML JS
    //1) Prerequisite
    <script src="/static/assets/socket.io-client/socket.io.js"></script>

    //2) Initialize io
    var socket = io.connect('http://localhost:8080/');
    
    //Send Something from client
    socket.emit("chat message", msg);

    //Get something to client
    socket.on("chat message",function(msg){
            console.log(msg);
    });
