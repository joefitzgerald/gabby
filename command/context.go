package command

type Context struct {
	TenantID       string `default:"common" hidden:"" env:"GABBY_TENANT_ID"`
	ClientID       string `default:"45c7f99c-0a94-42ff-a6d8-a8d657229e8c" hidden:"" env:"GABBY_CLIENT_ID"`
	TokenCachePath string `default:"token_cache.json" hidden:"" env:"GABBY_TOKEN_CACHE_PATH"`
	Me             string `-:"" hidden:""`
}
