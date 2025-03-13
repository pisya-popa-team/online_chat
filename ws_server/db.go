package wsserver

import "online_chat/database"

var (
	db = database.GetDBConnection()
)