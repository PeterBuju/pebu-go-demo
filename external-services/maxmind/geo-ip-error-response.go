package maxmind

type GeoIpErrorResponse struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}
