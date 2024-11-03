package handlers

import "online_chat/enviroment"

var (
	access_secret, refresh_secret = enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET"), enviroment.GoDotEnvVariable("REFRESH_TOKEN_SECRET")
)