package config

type Config struct {
	Mode string `toml:"mode"` // gin 模式
	Port string `toml:"port"` // api port
	Log  log    `toml:"log"`  // log
	DB   db     `toml:"db"`   //数据库

	Export string `toml:"export"` // 导出路径
}

//日志部分
type log struct {
	Level string `toml:"level"` // 日志级别
	Path  string `toml:"patyh"` // 日志地址
}

//etcd部分
type etcd struct {
	Endpoints []string `toml:"endpoints"` // etcd的集群列表, 配置多个, 避免单点故障
	Timeout   int      `toml:"timeout"`   // etcd的连接超时, 单位毫秒
	Username  string   `toml:"username"`  // 用户名
	Password  string   `toml:"password"`  // 密码
}

// DB
type db struct {
	Dialect         string `toml:"dialect"`           // 数据库类型, Default: mysql. Support soon: postgres,sqllite
	Host            string `toml:"host"`              // 数据库地址, localhost
	Port            string `toml:"port"`              // 数据库端口号 3306
	Name            string `toml:"name"`              // 数据库名 dev
	MaxIdleConns    int    `toml:"max_idle_conns"`    // 连接池中最大空闲连接数
	MaxOpenConns    int    `toml:"max_open_conns"`    // 打开的最大连接数
	ConnMaxLifetime int    `toml:"conn_max_lifetime"` // 连接的最大空闲时间, 单位秒
	Username        string `toml:"username"`          // 数据库用户名
	Password        string `toml:"password"`          // 数据库密码
}
