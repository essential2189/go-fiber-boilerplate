package config

type Config struct {
	DB     DB     `mapstructure:"db"`
	Redis  Redis  `mapstructure:"redis"`
	Kafka  Kafka  `mapstructure:"kafka"`
	Sentry Sentry `mapstructure:"sentry"`
}

type DB struct {
	User         string   `mapstructure:"user"`
	Password     string   `mapstructure:"password"`
	SourceAddrs  []string `mapstructure:"source_addrs"`
	ReplicaAddrs []string `mapstructure:"replica_addrs"`
	DBName       string   `mapstructure:"dbname"`
}

type Redis struct {
	MasterName    string   `mapstructure:"master_name"`
	SentinelAddrs []string `mapstructure:"sentiel_addrs"`
}
type Kafka struct {
	Version string   `mapstructure:"version"`
	Brokers []string `mapstructure:"brokers"`
	Group   string   `mapstructure:"group"`
	Topics  []string `mapstructure:"topics"`
}

type Sentry struct {
	Dsn string `mapstructure:"dsn"`
}
