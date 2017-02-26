package models

type Host struct {
	Ip string
	Port string
	PrivateKeyFilePath string `mapstructure:"private_key_file_path"`
	User string
}
