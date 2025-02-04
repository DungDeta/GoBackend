package setting

type Config struct {
	Server Server        `mapstructure:"server"`
	Mysql  MySqlSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
	Redis  RedisSetting  `mapstructure:"redis"`
}
type Server struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}
type MySqlSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DbName          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}
type LoggerSetting struct {
	LogLevel    string `mapstructure:"log_level"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	MaxSize     int    `mapstructure:"max_size"`
	Compress    bool   `mapstructure:"compress"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	Pool     int    `mapstructure:"poolSize"`
}
