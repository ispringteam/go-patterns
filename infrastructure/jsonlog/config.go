package jsonlog

// Config is used for configuration of logger
type Config struct {
	Level   Level  // Affects logging level
	AppName string // Is represented as field "app_name" in logs
}
