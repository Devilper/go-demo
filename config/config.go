package config

type LocalConfig struct {
	Db MysqlConfig `mapstructure:"db" json:"db"`
}

type MysqlConfig struct {
	DriverName string `mapstructure:"driver_name" json:"driver_name"`
	Host       string `mapstructure:"host" json:"host"`
	Port       int    `mapstructure:"port" json:"port"`
	Database   string `mapstructure:"database" json:"database"`
	Username   string `mapstructure:"username" json:"username"`
	Password   string `mapstructure:"password" json:"password"`
	Charset    string `mapstructure:"charset" json:"charset"`
}
