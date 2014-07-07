package cloudflare

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

const (
	CF_API_URL string = "https://www.cloudflare.com/api_json.html"
)

type CloudGuard struct {
	ApiKey string
	Email  string
	Domain string
	Debug  bool
}

func Connect(apikey, email, domain string, debug bool) *CloudGuard {
	return &CloudGuard{
		ApiKey: apikey,
		Email:  email,
		Domain: domain,
		Debug:  debug,
	}
}

func (this *CloudGuard) SendRequest(values url.Values, structure interface{}) (*interface{}, error) {
	values.Set("tkn", this.ApiKey)
	values.Set("email", this.Email)

	response, err := http.PostForm(CF_API_URL, values)
	if err != nil {
		log.Println(err)
		return err
	}
	defer req.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
	}

	data := Root{}
	data.Result = structure

	err = json.Unmarshal(content, &data)

	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (this *CloudGuard) GetDomainStats(interval string) (*Stats, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "stats")
	values.Set("interval", interval)

	data, err := this.SendRequest(values, Stats{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) GetDomainsList() (*ZonesLoad, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "zone_load_multi")

	data, err := this.SendRequest(values, ZonesLoad{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) GetDnsRecords() (*DnsRecords, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "rec_load_all")

	data, err := this.SendRequest(values, DnsRecords{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) NewDnsRecord(values ...map[string]string) (*Record, error) {
	args := url.Values
	args.Set("z", "rec_new")
	for key, value := range values {
		args.Set(key, value)
	}
	data, err := this.SendRequest(args, Record{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) EditDnsRecord(id string, values ...map[string]string) (*Record, error) {
	args := url.Values
	args.Set("z", "rec_edit")
	args.Set("id", id)
	for key, value := range values {
		args.Set(key, value)
	}
	data, err := this.SendRequest(args, Record{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) DeleteDnsRecord(id string) error {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("id", id)
	data, err := this.SendRequest(values, interface{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) SetProxyStatus(id string, status bool) error {
	proxy_status := "0"
	if status {
		proxy_status = "1"
	}
	records, err := this.GetDnsRecords()
	if err != nil {
		return err
	}
	for _, record := range records {
		if record.Id == id {
			values := make(map[string]string)
			values["service_mode"] = proxy_status
			_, err := this.EditDnsRecord(id, values)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("Unknow dns record id")
}

func (this *CloudGuard) SetSecurityLevel(level string) (*ZoneLoad, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "sec_lvl")
	values.Set("v", level)

	data, err := this.SendRequest(values, ZonesLoad{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) SetCacheLevel(level string) (*ZoneLoad, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "cache_lvl")
	values.Set("v", level)

	data, err := this.SendRequest(values, ZonesLoad{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) SetDevMode(enable bool) (*ZoneLoad, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "devmode")
	values.Set("v", level)

	data, err := this.SendRequest(values, ZonesLoad{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) PurgeCache() (*PurgeCache, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "fpurge_ts")
	values.Set("v", "1")

	data, err := this.SendRequest(values, PurgeCache{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) PurgeFile(url string) (*PurgeFile, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "fpurge_ts")
	values.Set("url", url)

	data, err := this.SendRequest(values, PurgeFile{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) LookupIp(ip string) error {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "ip_lkup")
	values.Set("ip", ip)

	_, err := this.SendRequest(values, interface{})
	if err != nil {
		return err
	}
	return nil
}

func (this *CloudGuard) DenyIP(ip string) (*ModIp, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "ban")
	values.Set("key", ip)

	data, err := this.SendRequest(values, ModIp{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) ForgetIP(ip string) (*ModIp, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "nul")
	values.Set("key", ip)

	data, err := this.SendRequest(values, ModIp{})
	if err != nil {
		return nil, err
	}
}

func (this *CloudGuard) AllowIP(ip string) (*ModIp, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "wl")
	values.Set("key", ip)

	data, err := this.SendRequest(values, ModIp{})
	if err != nil {
		return nil, err
	}
}

func (this *CloudGuard) ToggleMirage2(toggle bool) error {
	status := "0"
	if toggle {
		status = "1"
	}
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "mirage2")
	values.Set("v", status)

	_, err := this.SendRequest(values, interface{})
	if err != nil {
		return err
	}
	return nil
}

func (this *CloudGuard) Minify(state string) error {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "minify")
	values.Set("v", state)

	_, err := this.SendRequest(values, interface{})
	if err != nil {
		return err
	}
	return nil
}

func (this *CloudGuard) SetRocketLoader(state string) error {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "async")
	values.Set("v", state)

	_, err := this.SendRequest(values, interface{})
	if err != nil {
		return err
	}
	return nil
}

func (this *CloudGuard) ToggleIpv46(toggle bool) error {
	status := "0"
	if toggle {
		status = "3"
	}
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "ipv46")
	values.Set("v", status)

	_, err := this.SendRequest(values, interface{})
	if err != nil {
		return err
	}
	return nil
}

func (this *CloudGuard) Snapshot(zoneid string) error {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "zone_grab")
	values.Set("zid", zoneid)

	_, err := this.SendRequest(values, interface{})
	if err != nil {
		return err
	}
	return nil
}

func (this *CloudGuard) GetZoneSettings() (*ZoneSettings, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "zone_settings")

	data, err := this.SendRequest(values, ZoneSettings{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) GetActiveZones(zones ...string) (*ZonesCheck, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "zone_check")
	values.Set("zones", strings.Join(zones, ","))

	data, err := this.SendRequest(values, ZonesCheck{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (this *CloudGuard) GetRecentIps(hours, class, geo string) (*ZoneIps, error) {
	values := url.Values
	values.Set("z", this.Domain)
	values.Set("a", "zone_ips")
	values.Set("hours", hours)
	values.Set("class", class)
	values.Set("geo", geo)

	data, err := this.SendRequest(values, ZoneIps{})
	if err != nil {
		return nil, err
	}
	return &data, nil
}
