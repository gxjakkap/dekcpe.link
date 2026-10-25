package utils

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetIPFromHeaders(t *testing.T) {
	cases := []struct {
		mode, header, value, want string
	}{
		{"xff", fiber.HeaderXForwardedFor, "6.6.6.6, 1.2.3.4", "1.2.3.4"},
		{"xff", fiber.HeaderXForwardedFor, "1.2.3.4", "1.2.3.4"},
		{"xff", fiber.HeaderXForwardedFor, "1.2.3.4, not-an-ip", "0.0.0.0"},
		{"xff", fiber.HeaderXForwardedFor, "", "0.0.0.0"},
		{"cf_argo", "CF-Connecting-IP", "1.2.3.4", "1.2.3.4"},
		{"cf_argo", "CF-Connecting-IP", "../evil?x=", "0.0.0.0"},
	}

	for _, tc := range cases {
		t.Setenv("PROXY_MODE", tc.mode)
		app := fiber.New()
		app.Get("/", func(c *fiber.Ctx) error {
			ip, _ := GetIPFromHeaders(c)
			return c.SendString(ip)
		})

		req := httptest.NewRequest("GET", "/", nil)
		if tc.value != "" {
			req.Header.Set(tc.header, tc.value)
		}
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}
		got, _ := io.ReadAll(resp.Body)
		if string(got) != tc.want {
			t.Errorf("%s %q=%q: got %q, want %q", tc.mode, tc.header, tc.value, got, tc.want)
		}
	}
}
