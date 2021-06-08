package conf

type Server struct {
	Database Database `yaml:"Database"`
	Sys Sys	`yaml:"Sys"`
	Jwt Jwt `yaml:"jwt"`
}

type Database struct {
	DBType       string `yaml:"DBType"`
	UserName     string	`yaml:"UserName"`
	Password     string	`yaml:"Password"`
	Host         string	`yaml:"Host"`
	DBName       string	`yaml:"DBName"`
	Charset      string	`yaml:"Charset"`
	ParseTime    bool	`yaml:"ParseTime"`
	MaxIdleConns int	`yaml:"MaxIdleConns"`
	MaxOpenConns int	`yaml:"MaxOpenConns"`
}

type Sys struct {
	RunMode string `yaml:"RunMode"`
	HttpPort string	`yaml:"HttpPort"`
	ReadTimeout int64	`yaml:"ReadTimeout"`
	WriteTimeout int64	`yaml:"WriteTimeout"`
}

type Jwt struct {
	SignKey string `yaml:"signkey"`
}