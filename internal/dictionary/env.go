package dictionary

import (
	"os"
)

type allEnv struct {
	ListenedAddress string
	DbHost          string
	DbName          string
	DbUser          string
	DbPassword      string
}

// Env Получаем переменные окружения
var Env = &allEnv{
	ListenedAddress: os.Getenv("LISTENED_ADDRESS"),
	DbHost:          os.Getenv("DB_HOST"),
	DbName:          os.Getenv("DB_NAME"),
	DbUser:          os.Getenv("DB_USER"),
	DbPassword:      os.Getenv("DB_PASSWORD"),
}
