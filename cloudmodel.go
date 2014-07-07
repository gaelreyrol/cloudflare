package cloudflare

const (
	E_UNAUTH     string = "E_UNAUTH"
	E_INVLDINPUT string = "E_INVLDINPUT"
	E_MAXAPI     string = "E_MAXAPI"
)

type Root struct {
	Response string      `json:"response"`
	Result   interface{} `json:"result"`
	Message  string      `json:"msg"`
}

/* Strutures for Stats Request */

type Stats struct {
	TimeZero float64      `json:"timeZero"`
	TimeEnd  float64      `json:"timeEnd"`
	Count    int          `json:"count"`
	HasMore  bool         `json:"has_more"`
	Objs     []StatsChild `json:"objs"`
}

type StatsChild struct {
	CachedServerTime    float64        `json:"cachedServerTime"`
	CachedExpryTime     float64        `json:"cachedExpryTime"`
	TrafficBreakdown    []TrafficChild `json:"trafficBreakdown"`
	BandwidthServed     ServedStats    `json:"bandwidthServed"`
	requestsServed      ServedStats    `json:"requestsServed"`
	ProZone             bool           `json:"pro_zone"`
	PageLoadTime        string         `json:"pageLoadTime"`
	CurrentServerTime   float64        `json:"currentServerTime"`
	Interval            int            `json:"interval"`
	ZoneCDate           float64        `json:"zoneCDate"`
	UserSecuritySetting string         `json:"userSecuritySetting"`
	DevMode             int            `json:dev_mode`
	Ipv46               int            `json:"ipv46"`
	Ob                  int            `json:"op"`
	CacheLevel          string         `json:"cache_lvl"`
}

type TrafficChild struct {
	Pageviews TrafficChildStats `json:"pagesviews"`
	Uniques   TrafficChildStats `json:"uniques"`
}

type TrafficChildStats struct {
	Regular int `json:"regular"`
	Threat  int `json:"threat"`
	Crawler int `json:"crawler"`
}

type ServedStats struct {
	Cloudflare float64 `json:"cloudflare"`
	User       float64 `json:"user"`
}

/* End Strutures for Stats Request */

/* Strutures for Zone Multi Load Request */

type ZonesLoad struct {
	HasMore bool       `json:"has_more"`
	Count   int        `json:"count"`
	Objs    []ZoneLoad `json:"objs"`
}

type ZoneLoad struct {
	ZoneId          string           `json:"zone_id"`
	UserId          string           `json:"user_id"`
	ZoneName        string           `json:"zone_name"`
	DisplayName     string           `json:"display_name"`
	ZoneStatus      string           `json:"zone_status"`
	ZoneMode        string           `json:"zone_mode"`
	HostId          string           `json:"host_id"`
	ZoneType        string           `json:"zone_type"`
	HostPubName     string           `json:"host_pubname"`
	HostWebsite     string           `json:"host_website"`
	Vtxt            string           `json:"vtxt"`
	Fqdns           []string         `json:"fqdns"`
	Step            string           `json:"step"`
	ZoneStatusClass string           `json:"zone_status_class"`
	ZoneStatusDesc  string           `json:"zone_status_desc"`
	NsVanityMap     []interface{}    `json:"ns_vanity_map"`
	OrigRegistrar   string           `json:"orig_registrar"`
	OrigDnshost     string           `json:"orig_dnshost"`
	OrigNsnames     string           `json:"orig_dnshost"`
	Props           ZoneLoadProperty `json:"props"`
	ConfirmCode     []Codes          `json:"confirm_code"`
	Allow           []string         `json:"allow"`
}

type ZoneLoadProperty struct {
	DnsCName       int `json:"dns_cname"`
	DnsPartner     int `json:"dns_partner"`
	DnsAnonPartner int `json:"dns_anon_partner"`
	Pro            int `json:"pro"`
	ExpiredPro     int `json:"exprired_pro"`
	ProSub         int `json:"pro_sub"`
	Ssl            int `json:"ssl"`
	ExpiredSsl     int `json:"expired_ssl"`
	ExpriredRsPro  int `json:"expired_rs_pro"`
	ResellerPro    int `reseller_pro`
	ForceInteral   int `json:"force_interal"`
	SslNeeded      int `json:"ssl_needed"`
	AlexaRank      int `json:"alexa_rank"`
}

type Codes struct {
	ZoneDeactivate string `json:"zone_deactivate"`
	ZoneDevModel   string `json:"zone_dev_mode1"`
}

/* End Strutures for Zone Multi Load Request */

/* Strutures for Dns Records Request */

type DnsRecords struct {
	HasMore bool     `json:"has_more"`
	Count   int      `json:"count"`
	Objs    []Record `json:"objs"`
}

type Record struct {
	Id             string      `json:"rec_id"`
	Tag            string      `json:"rec_tag"`
	ZoneName       string      `json:"zone_name"`
	Name           string      `json:"name"`
	DisplayName    string      `json:"display_name"`
	Type           string      `json:"type"`
	Prio           string      `json:"prio"`
	Content        string      `json:"content"`
	DisplayContent string      `json:"display_content"`
	Ttl            string      `json:"ttl"`
	TtlCeil        int         `json:"ttl_ceil"`
	SslId          string      `json:"ssl_id"`
	SslStatus      string      `json:"ssl_status"`
	SslExpiresOn   string      `json:"ssl_expires_on"`
	AutoTtl        int         `json:"auto_ttl"`
	ServiceMode    string      `json:"service_mode"`
	Props          DnsProperty `json:"props"`
}

type DnsProperty struct {
	Proxiable   int `json:"proxiable"`
	CloudOn     int `json:"cloud_on"`
	CfOpen      int `json:"cf_open"`
	Ssl         int `json:"ssl"`
	ExpiredSsl  int `json:"expired_ssl"`
	ExpiringSsl int `json:"expiring_ssl"`
	PendingSsl  int `json:"pending_ssl"`
}

/* End Strutures for Dns Records Request */

/* Strutures for Zones Check Request */

type ZonesCheck struct {
	Zones map[string]int `json:"zones"`
}

/* END Strutures for Zones Check Request */

/* Strutures for Zones Ips Request */

type ZoneIps struct {
	Ips Ip `json:"ips"`
}

type Ip struct {
	Ip             string  `json:"ip"`
	Classification string  `json:"classification"`
	Hits           string  `json:"hits"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	ZoneName       string  `json:"zone_name"`
}

/* END Strutures for Zones Ips Request */

/* Strutures for Zone Settings Request */

type ZoneSettings struct {
	Result []Settings `json:"result"`
}

type Settings struct {
	UserSecuritySetting string `json:"userSecuritySetting"`
	DevMode             int    `json:"dev_mode"`
	Ipv46               int    `json:"ipv46"`
	Ob                  int    `json:"ob"`
	CacheLevel          string `json:"cache_lvl"`
	OutboundLinks       string `json:"outboundLinks"`
	Async               string `json:"async"`
	Bic                 string `json:"bic"`
	ChlTtl              string `json:"chl_ttl"`
	ExpTtl              string `json:"exp_ttl"`
	FpurgeTs            string `json:"fpurge_ts"`
	Hotling             string `json:"hotlink"`
	Img                 string `json:"img"`
	Lazy                string `json:"lazy"`
	Minify              string `json:"minify"`
	Outlink             string `json:"outlink"`
	Preload             string `json:"preload"`
	S404                string `json:"s404"`
	SecLvl              string `json:"sec_lvl"`
	Sdpy                string `json:"sdpy"`
	Ssl                 string `json:"ssl"`
	WafProfile          string `json:"waf_profile"`
}

/* END Strutures for Zone Settings Request */

type SecLevel struct {
	Zone ZoneLoad `json:"zone"`
}

type CacheLevel struct {
	Zone ZoneLoad `json:"zone"`
}

type DevMode struct {
	ExpiresOn float64  `json:"expires_on"`
	Zone      ZoneLoad `json:"zone"`
}

type PurgeCache struct {
	FpurgeTs float64  `json:"fpurge_ts"`
	Zone     ZoneLoad `json:"zone"`
}

type PurgeFile struct {
	Vtxt_match string `json:"vtxt_match"`
	Url        string `json:"url"`
}

type ModIp struct {
	Ip     string `json:"ip"`
	Action string `json:"action"`
}

type NewRecord struct {
	Rec Record `json:"obj"`
}

type EditRecord struct {
	Rec Record `json:"obj"`
}
