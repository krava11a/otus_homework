package models

type App struct {
	ID     int    `yaml:"id"`
	Name   string `yaml:"name"`
	Secret string `yaml:"secret"`
}
