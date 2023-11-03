package config

type Config struct {
	Log    Log    `mapstructure:"log"`
	DB     DB     `mapstructure:"db"`
	Kafka  Kafka  `mapstructure:"kafka"`
	Redis  Redis  `mapstructure:"kafka"`
	Sentry Sentry `mapstructure:"sentry"`
}

type Log struct {
	HistoryType string `mapstructure:"history_type"`
	Version     string `mapstructure:"version"`
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
	Broker []string `mapstructure:"broker"`
}

type Sentry struct {
	Dsn string `mapstructure:"dsn"`
}
