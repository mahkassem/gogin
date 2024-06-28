package config

type Response struct {
	Message string
	Data    interface{}
	Error   *Error
}

type Error struct {
	Error    error
	Message  string
	Code     int
	MetaData interface{}
}
