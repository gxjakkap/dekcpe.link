package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gxjakkap/dekcpe.link/model"
)

type ipApiResp struct {
	Query       string  `json:"query"`
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
}

var geoClient = &http.Client{Timeout: 5 * time.Second}

func GetGeoFromIP(ip string) (*model.GeoLocation, error) {
	resp, err := geoClient.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
	if err != nil {
		return nil, fmt.Errorf("geolocation request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("geolocation request: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("geolocation read body: %w", err)
	}

	d := ipApiResp{}
	if err := json.Unmarshal(body, &d); err != nil {
		return nil, fmt.Errorf("geolocation parse body: %w", err)
	}

	if d.Status != "success" {
		return nil, fmt.Errorf("geolocation lookup status %q", d.Status)
	}

	r := &model.GeoLocation{
		Country:  d.Country,
		Region:   d.RegionName,
		City:     d.City,
		Timezone: d.Timezone,
	}

	return r, nil
}

func GetIPFromHeaders(c *fiber.Ctx) (ip string, err error) {
	pm := os.Getenv("PROXY_MODE")

	if pm == "" {
		return c.IP(), nil
	}

	if pm == "cf_argo" {
		cfIP := strings.TrimSpace(c.Get("CF-Connecting-IP"))
		if net.ParseIP(cfIP) == nil {
			return c.IP(), nil
		}
		return cfIP, nil
	}

	// Use the rightmost entry: it's the one our proxy appended. Everything to
	// its left is client-supplied and spoofable.
	// ponytail: assumes exactly one proxy hop; with a proxy chain, skip one entry per trusted hop from the right.
	ips := strings.Split(c.Get(fiber.HeaderXForwardedFor), ",")
	last := strings.TrimSpace(ips[len(ips)-1])
	if net.ParseIP(last) == nil {
		return c.IP(), nil
	}

	return last, nil
}

func GetForwardingHeader() (header string) {
	pm := os.Getenv("PROXY_MODE")
	log.Printf("[DEBUG] Proxy Mode: %s", pm)

	if pm == "" {
		return ""
	}

	if pm == "cf_argo" {
		log.Println("[DEBUG] Setting Forwarding Header as CF-Connecting-IP")
		return "CF-Connecting-IP"
	}

	log.Println("[DEBUG] Setting Forwarding Header as X-Forwarded-For")

	return fiber.HeaderXForwardedFor
}
