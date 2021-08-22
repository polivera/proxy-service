package src

const (
	CacheStatusDoNotSave   = 0
	CacheStatusDoNotUse    = 1
	CacheStatusReturnCache = 2

	// DefaultRowsPerPage This should be a string because it is send on GetQueryParam
	DefaultRowsPerPage        = "5"
	PagePaginationQueryParam  = "page"
	LimitPaginationQueryParam = "rows"
	RequestURLQueryParam      = "request_url"
	RequestMethodQueryParam   = "request_method"
)
