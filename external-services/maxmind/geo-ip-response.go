package maxmind

// GeoIpResponse represents the structure of the response from the Geo IP lookup
type GeoIpResponse struct {
	Continent         Continent         `json:"continent"`
	Country           Country           `json:"country"`
	Location          Location          `json:"location"`
	RegisteredCountry RegisteredCountry `json:"registered_country"`
	Traits            Traits            `json:"traits"`
}

type Continent struct {
	Code      string            `json:"code"`
	GeonameID int               `json:"geoname_id"`
	Names     map[string]string `json:"names"`
}

type Country struct {
	GeonameID int               `json:"geoname_id"`
	IsoCode   string            `json:"iso_code"`
	Names     map[string]string `json:"names"`
}

type Location struct {
	AccuracyRadius int     `json:"accuracy_radius"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	TimeZone       string  `json:"time_zone"`
}

type RegisteredCountry struct {
	GeonameID int               `json:"geoname_id"`
	IsoCode   string            `json:"iso_code"`
	Names     map[string]string `json:"names"`
}

type Traits struct {
	AutonomousSystemNumber       int    `json:"autonomous_system_number"`
	AutonomousSystemOrganization string `json:"autonomous_system_organization"`
	ConnectionType               string `json:"connection_type"`
	IPAddress                    string `json:"ip_address"`
	ISP                          string `json:"isp"`
	Network                      string `json:"network"`
	Organization                 string `json:"organization"`
}
