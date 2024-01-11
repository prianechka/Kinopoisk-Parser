package config

type ServerConfig struct {
	StartPort    string                   `toml:"start_port"`
	DbParams     DatabaseConnectionParams `toml:"database"`
	MaxMovies    uint64                   `toml:"max_movies"`
	TimeForSleep uint64                   `toml:"time_for_sleep"`
	MovieURL     string                   `toml:"movie_url"`
	PersonURL    string                   `toml:"person_url"`
	Token        string                   `toml:"token"`
}

type DatabaseConnectionParams struct {
	Scheme   string `toml:"scheme"`
	Host     string `toml:"host"`
	Port     uint64 `toml:"port"`
	Database string `toml:"database"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

func CreateConfig() *ServerConfig {
	return &ServerConfig{}
}
