package models

// ProxyRequest Save all request that pass through the proxy
type ProxyRequest struct {
	ID              string
	Method          string
	Host            string
	URI             string
	Headers         string
	Body            string
	ResponseHeaders string
	ResponseBody    string
	ResponseStatus  uint
	Status          uint8
}

// RequestConfig Save where to proxy the incoming requests
type RequestConfig struct {
	ID            uint
	Source        string
	Target        string
	OriginHeader  string
	RefererHeader string
}
