package config

type Config struct {
	Kafka     Kafka     `yaml:"kafka"`
	Api       Api       `yaml:"api"`
	Inference Inference `yaml:"inference"`
}

type Kafka struct {
	Servers                []string `yaml:"servers"`
	RedditCommentsTopic    string   `yaml:"redditCommentsTopic"`
	RedditSubmissionsTopic string   `yaml:"redditSubmissionsTopic"`
}

type Api struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Inference struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
