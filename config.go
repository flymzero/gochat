package gochat

//--------------------------------------common-------------------------------------------
const (
	TEXT_INTO_ROOM  = " [%s] into room "
	TEXT_LEAVE_ROOM = " [%s] leave room "
)

//--------------------------------------server-------------------------------------------
const (
	CONFIG_SERVER_PROTOCOL = "tcp"
	CONFIG_SERVER_IP       = "127.0.0.1"
	CONFIG_SERVER_PORT     = ":9000"
)

//--------------------------------------client-------------------------------------------
const (
	//check client is alive (10 minutes)
	CONFIG_CLIENT_TIMEOUT = 600
)

//----------------------------------------gui-------------------------------------------
