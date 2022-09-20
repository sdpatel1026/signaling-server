package configs

const (
	ENV_FILE     string = ".env"
	KEY_APP_PORT string = "APP_PORT"
	KEY_ROOM_ID  string = "room_id"
	KEY_ERROR    string = "error"

	ERROR_EMPTY_ROOM_ID string = "room_id is missing in url params."
	SERVER_ERROR        string = "Something went wrong at the server side. Please, try again later."
)
