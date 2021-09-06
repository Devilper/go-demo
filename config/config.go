package config

type LocalConfig struct {
	Db    MySqlConfig `mapstructure:"db" json:"db"`
	Redis RedisConfig `mapstructure:"redis" json:"redis"`
}

type MySqlConfig struct {
	DriverName string `mapstructure:"driver_name" json:"driver_name"`
	Host       string `mapstructure:"host" json:"host"`
	Port       string `mapstructure:"port" json:"port"`
	Database   string `mapstructure:"database" json:"database"`
	Username   string `mapstructure:"username" json:"username"`
	Password   string `mapstructure:"password" json:"password"`
	Charset    string `mapstructure:"charset" json:"charset"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address" json:"address"`
	Password string `mapstructure:"password" json:"password"`
	Db       int    `mapstructure:"db" json:"db"`
}
