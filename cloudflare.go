package cloudflare

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	CF_API_URL string = "https://www.cloudflare.com/api_json.html"
)

type Cloudflare struct {
	ApiKey string
	Email  string
	Domain string
	Debug  bool
}

func Connect(apikey, email string, debug bool) *Cloudflare {
	return &Cloudflare{
		ApiKey: apikey,
		Email:  email,
		Debug:  debug,
	}
}

func (this *Cloudflare) sendRequest(values url.Values) ([]byte, error) {
	values.Set("tkn", this.ApiKey)
	values.Set("email", this.Email)

	response, err := http.PostForm(CF_API_URL, values)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return content, nil
}

func (this *Cloudflare) GetDomainStats(domain, interval string) (RootStats, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "stats")
	values.Set("interval", interval)

	response, err := this.sendRequest(values)

	data := RootStats{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootStats{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootStats{}, err
	}
	return data, nil
}

func (this *Cloudflare) GetDomainsList() (RootZones, error) {
	values := url.Values{}
	values.Set("a", "zone_load_multi")

	response, err := this.sendRequest(values)

	data := RootZones{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootZones{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootZones{}, err
	}
	return data, nil
}

func (this *Cloudflare) GetDnsRecords(domain string) (RootDnsRecords, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "rec_load_all")

	response, err := this.sendRequest(values)

	data := RootDnsRecords{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootDnsRecords{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootDnsRecords{}, err
	}
	return data, nil
}

func (this *Cloudflare) NewDnsRecord(domain string, values map[string]string) (RootNewRecord, error) {
	args := url.Values{}
	args.Set("z", domain)
	args.Set("a", "rec_new")
	for k, _ := range values {
		args.Set(k, values[k])
	}

	response, err := this.sendRequest(args)

	data := RootNewRecord{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootNewRecord{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootNewRecord{}, err
	}
	return data, nil
}

func (this *Cloudflare) EditDnsRecord(domain, id string, values map[string]string) (RootEditRecord, error) {
	args := url.Values{}
	args.Set("z", domain)
	args.Set("a", "rec_edit")
	args.Set("id", id)
	for k, _ := range values {
		args.Set(k, values[k])
	}

	response, err := this.sendRequest(args)

	data := RootEditRecord{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootEditRecord{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootEditRecord{}, err
	}
	return data, nil
}

func (this *Cloudflare) DeleteDnsRecord(domain, id string) (Root, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("id", id)

	response, err := this.sendRequest(values)

	data := Root{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return Root{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return Root{}, err
	}
	return data, nil
}

func (this *Cloudflare) SetProxyStatus(domain, id string, status bool) (RootEditRecord, error) {
	proxy_status := "0"
	if status {
		proxy_status = "1"
	}

	values := make(map[string]string)
	values["service_mode"] = proxy_status

	response, err := this.EditDnsRecord(domain, id, values)

	return response, err
}

func (this *Cloudflare) SetSecurityLevel(domain, level string) (RootZones, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "sec_lvl")
	values.Set("v", level)

	response, err := this.sendRequest(values)

	data := RootZones{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootZones{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootZones{}, err
	}
	return data, nil
}

func (this *Cloudflare) SetCacheLevel(domain, level string) (RootZones, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "cache_lvl")
	values.Set("v", level)

	response, err := this.sendRequest(values)

	data := RootZones{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootZones{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootZones{}, err
	}
	return data, nil
}

func (this *Cloudflare) SetDevMode(domain string, enable bool) (RootZones, error) {
	dev := "0"
	if enable {
		dev = "1"
	}

	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "devmode")
	values.Set("v", dev)

	response, err := this.sendRequest(values)

	data := RootZones{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootZones{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootZones{}, err
	}
	return data, nil
}

func (this *Cloudflare) PurgeCache(domain string) (RootPurgeCache, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "fpurge_ts")
	values.Set("v", "1")

	response, err := this.sendRequest(values)

	data := RootPurgeCache{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootPurgeCache{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootPurgeCache{}, err
	}
	return data, nil
}

func (this *Cloudflare) PurgeFile(domain, url_file string) (RootPurgeFile, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "zone_file_purge")
	values.Set("url", url_file)

	response, err := this.sendRequest(values)

	data := RootPurgeFile{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootPurgeFile{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootPurgeFile{}, err
	}
	return data, nil
}

func (this *Cloudflare) LookupIp(domain, ip string) (RootLookupIp, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "ip_lkup")
	values.Set("ip", ip)

	response, err := this.sendRequest(values)

	data := RootLookupIp{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootLookupIp{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootLookupIp{}, err
	}
	return data, nil
}

func (this *Cloudflare) DenyIP(domain, ip string) (RootModIp, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "ban")
	values.Set("key", ip)

	response, err := this.sendRequest(values)

	data := RootModIp{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootModIp{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootModIp{}, err
	}
	return data, nil
}

func (this *Cloudflare) ForgetIP(domain, ip string) (RootModIp, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "nul")
	values.Set("key", ip)

	response, err := this.sendRequest(values)

	data := RootModIp{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootModIp{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootModIp{}, err
	}
	return data, nil
}

func (this *Cloudflare) AllowIP(domain, ip string) (RootModIp, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "wl")
	values.Set("key", ip)

	response, err := this.sendRequest(values)

	data := RootModIp{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootModIp{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootModIp{}, err
	}
	return data, nil
}

func (this *Cloudflare) ToggleMirage2(domain string, toggle bool) (Root, error) {
	status := "0"
	if toggle {
		status = "1"
	}
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "mirage2")
	values.Set("v", status)

	response, err := this.sendRequest(values)

	data := Root{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return Root{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return Root{}, err
	}
	return data, nil
}

func (this *Cloudflare) Minify(domain, state string) (Root, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "minify")
	values.Set("v", state)

	response, err := this.sendRequest(values)

	data := Root{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return Root{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return Root{}, err
	}
	return data, nil
}

func (this *Cloudflare) SetRocketLoader(domain, state string) (Root, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "async")
	values.Set("v", state)

	response, err := this.sendRequest(values)

	data := Root{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return Root{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return Root{}, err
	}
	return data, nil
}

func (this *Cloudflare) ToggleIpv46(domain string, toggle bool) (Root, error) {
	status := "0"
	if toggle {
		status = "3"
	}
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "ipv46")
	values.Set("v", status)

	response, err := this.sendRequest(values)

	data := Root{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return Root{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return Root{}, err
	}
	return data, nil
}

func (this *Cloudflare) Snapshot(domain, zoneid string) (Root, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "zone_grab")
	values.Set("zid", zoneid)

	response, err := this.sendRequest(values)

	data := Root{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return Root{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return Root{}, err
	}
	return data, nil
}

func (this *Cloudflare) GetZoneSettings(domain string) (RootZoneSettings, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "zone_settings")

	response, err := this.sendRequest(values)

	data := RootZoneSettings{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootZoneSettings{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootZoneSettings{}, err
	}
	return data, nil
}

func (this *Cloudflare) GetActiveZones(domain string, zones ...string) (RootZonesCheck, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "zone_check")
	values.Set("zones", strings.Join(zones, ","))

	response, err := this.sendRequest(values)

	data := RootZonesCheck{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootZonesCheck{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootZonesCheck{}, err
	}
	return data, nil
}

func (this *Cloudflare) GetRecentIps(domain, hours, class, geo string) (RootZoneIps, error) {
	values := url.Values{}
	values.Set("z", domain)
	values.Set("a", "zone_ips")
	values.Set("hours", hours)
	values.Set("class", class)
	values.Set("geo", geo)

	response, err := this.sendRequest(values)

	data := RootZoneIps{}
	err = json.Unmarshal(response, &data)
	if err != nil {
		return RootZoneIps{}, err
	}
	if data.Result == "error" {
		err = errors.New(data.Message)
		return RootZoneIps{}, err
	}
	return data, nil
}
