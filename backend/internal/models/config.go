package models

type Config struct {
	Environments []Environment `json:"environments"`
}

type Environment struct {
	Env string `json:"env"`
	Title string `json:"title"`
	Host string `json:"host"`
	Port int `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type FrontendEnvironment struct {
	Env string `json:"env"`
	Title string `json:"title"`
}