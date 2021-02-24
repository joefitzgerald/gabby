package command

type Context struct {
	TenantID           string `default:"common" hidden:"" env:"GABBY_TENANT_ID"`
	ClientID           string `default:"77605d92-6a9c-471e-a2d3-48c909b19ec8" hidden:"" env:"GABBY_CLIENT_ID"`
	TokenCacheFilename string `default:"token_cache.json" hidden:"" env:"GABBY_TOKEN_CACHE_FILENAME"`
	Me                 string `-:"" hidden:""`
}
