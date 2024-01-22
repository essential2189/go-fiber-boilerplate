package config

type Config struct {
	Server Server `mapstructure:"server"`
	App    App    `mapstructure:"app"`
	Infra  Infra  `mapstructure:"infra"`
}

type Server struct {
	Profile string `mapstructure:"profile"`
	Port    string `mapstructure:"port"`
}

type App struct {
	Version string `mapstructure:"version"`

	Url Url `mapstructure:"url"`

	Log Log `mapstructure:"log"`
}

type Infra struct {
	DB     DB     `mapstructure:"db"`
	Redis  Redis  `mapstructure:"redis"`
	Kafka  Kafka  `mapstructure:"kafka"`
	Sentry Sentry `mapstructure:"sentry"`
}

type Url struct {
	Billing struct {
		BaseUrl string `mapstructure:"base-url"`
	} `mapstructure:"billing"`
}

type Log struct {
	Level       string `mapstructure:"level"`
	HistoryType string `mapstructure:"history-type"`
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
