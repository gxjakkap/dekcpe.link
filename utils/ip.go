package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

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

func GetGeoFromIP(ip string) (*model.GeoLocation, error) {
	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))

	if err != nil {
		log.Fatalln("Failed request for geolocation info")
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		log.Fatalln("Failed parsing body for geolocation info")
		return nil, err
	}

	d := ipApiResp{}

	err = json.Unmarshal(body, &d)

	if err != nil {
		log.Fatalln("Failed parsing body for geolocation info")
		return nil, err
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
		cfIP := c.Get("CF-Connecting-IP")
		if cfIP == "" {
			return c.IP(), nil
		}
		return cfIP, nil
	}

	forwardedFor := c.Get(fiber.HeaderXForwardedFor)
	if forwardedFor == "" {
		return c.IP(), nil
	}

	ips := strings.Split(forwardedFor, ",")
	if len(ips) == 0 || strings.TrimSpace(ips[0]) == "" {
		return c.IP(), nil
	}

	return strings.TrimSpace(ips[0]), nil
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
