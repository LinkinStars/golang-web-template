package conf

// AllConfig 全部配置项
type AllConfig struct {
	Data   *Data `json:"data" mapstructure:"data"`
	Server *Server
}

type Server struct {
	HTTP HTTP `json:"http" mapstructure:"http"`
}

type HTTP struct {
	Addr string `json:"addr" mapstructure:"addr"`
}

type Data struct {
	Database Database `json:"database" mapstructure:"database"`
}

type Database struct {
	Connection      string `json:"connection" mapstructure:"connection"`
	ConnMaxLifeTime int    `json:"conn_max_life_time" mapstructure:"conn_max_life_time"`
	MaxOpenConn     int    `json:"max_open_conn" mapstructure:"max_open_conn"`
	MaxIdleConn     int    `json:"max_idle_conn" mapstructure:"max_idle_conn"`
}
