package config

type Config struct {
	Args map[string]interface{}
	Jobs []Job
}

type Job struct {
	Input  string
	Output string
}
