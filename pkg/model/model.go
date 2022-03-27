package model

// Config struct for yaml conf
type Config struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type Response struct {
	Action string  `json:"action"`
	Answer float64 `json:"answer"`
	Cached bool    `json:"cached"`
	X      float64 `json:"x" schema:"x"`
	Y      float64 `json:"y" schema:"y"`
}
