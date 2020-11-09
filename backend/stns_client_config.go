package backend

type STNSClientConfig struct {
	// api_endpoint = "http://api01.example.com/v1"
	ApiEndpoint string
	// http_proxy = "http://localhost:8080"
	HTTPProxy string
	// request_timeout = 3
	RequestTimeout int
	// request_retry = 1
	RequestRetry int
	// request_locktime = 600
	RequestLocktime int
	// ssl_verify = true
	SSLVerify bool
	// # basic auth
	// user = "basic_user"
	User string
	// password = "basic_password"
	Password string
	// # token auth
	// auth_token = "token"
	AuthToken string

	// query_wrapper = "/usr/local/bin/stns-query-wrapper"
	QueryWrapper string
	// chain_ssh_wrapper = "/usr/libexec/openssh/ssh-ldap-wrapper"
	ChainSSHWrapper string

	// cache = true
	Cache bool
	// cache_dir = "/var/cache/stns/"
	CacheDir string
	// cache_ttl = 600
	CacheTTL int
	// negative_cache_ttl = 600
	NegativeCacheTTL int

	// uid_shift = 2000
	UIDShift int
	// gid_shift = 2000
	GIDShift int

	// # tls client authentication
	// [tls]
	// ca   = "/etc/stns/keys/ca.pem"
	// cert = "/etc/stns/keys/client.crt"
	// key  = "/etc/stns/keys/client.key"
	TLS *TLS

	// [cached]
	// enable = true
	// prefetch = true
	Cached *Cached
}

type TLS struct {
	CA   string
	Cert string
	Key  string
}

type Cached struct {
	Enable   bool
	Prefecth bool
}
