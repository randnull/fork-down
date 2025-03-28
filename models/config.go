package models

type ConfigRepository struct {
	Host string
	Port string
}

type ConfigRestore struct {
	SaveFilePath string
}

type Config struct {
	ConfigRepository
	ConfigRestore
}
