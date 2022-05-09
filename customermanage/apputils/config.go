package apputils

type Config struct {
	Server   string
	LogLevel int
}

var AppConfig Config

func init() {
	AppConfig = Config{}
}
