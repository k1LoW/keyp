package backend

type STNSClientConfig struct {
	// api_endpoint = "http://api01.example.com/v1"
	APIEndpoint string `toml:"api_endpoint"`
	// http_proxy = "http://localhost:8080"
	HTTPProxy string `toml:"http_proxy"`
	// request_timeout = 3
	RequestTimeout int `toml:"request_timeout"`
	// request_retry = 1
	RequestRetry int `toml:"request_timeout"`
	// request_locktime = 600
	RequestLocktime int `toml:"request_locktime"`
	// ssl_verify = true
	SSLVerify bool `toml:"ssl_verify"`
	// # basic auth
	// user = "basic_user"
	User string `toml:"user"`
	// password = "basic_password"
	Password string `toml:"password"`
	// # token auth
	// auth_token = "token"
	AuthToken string `toml:"auth_token"`

	// query_wrapper = "/usr/local/bin/stns-query-wrapper"
	QueryWrapper string `toml:"query_wrapper"`
	// chain_ssh_wrapper = "/usr/libexec/openssh/ssh-ldap-wrapper"
	ChainSSHWrapper string `toml:"chain_ssh_wrapper"`

	// cache = true
	Cache bool `toml:"cache"`
	// cache_dir = "/var/cache/stns/"
	CacheDir string `toml:"cache_dir"`
	// cache_ttl = 600
	CacheTTL int `toml:"cache_ttl"`
	// negative_cache_ttl = 600
	NegativeCacheTTL int `toml:"negative_cache_ttl"`

	// uid_shift = 2000
	UIDShift int `toml:"uid_shift"`
	// gid_shift = 2000
	GIDShift int `toml:"gid_shift"`

	// # tls client authentication
	// [tls]
	// ca   = "/etc/stns/keys/ca.pem"
	// cert = "/etc/stns/keys/client.crt"
	// key  = "/etc/stns/keys/client.key"
	TLS *TLS `toml:"tls,omitempty"`

	// [cached]
	// enable = true
	// prefetch = true
	Cached *Cached `toml:"cached,omitempty"`
}

type TLS struct {
	CA   string `toml:"ca"`
	Cert string `toml:"cert"`
	Key  string `toml:"key"`
}

type Cached struct {
	Enable   bool `toml:"enable"`
	Prefecth bool `toml:"prefetch"`
}
