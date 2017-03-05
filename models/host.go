package models

// Host - struct for properties of a Host
type Host struct {
	IP                 string
	Port               string
	PrivateKeyFilePath string `mapstructure:"private_key_file_path"`
	User               string
}
