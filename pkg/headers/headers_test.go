package headers_test

import (
	"testing"

	"github.com/mallardduck/go-http-helpers/pkg/headers"
)

func TestAllGroupedHeaders(t *testing.T) {
	tests := []struct {
		name     string
		got      string
		expected string
	}{
		// Auth
		{"Auth.Authorization", headers.Auth.Authorization(), "Authorization"},
		{"Auth.ProxyAuthorization", headers.Auth.ProxyAuthorization(), "Proxy-Authorization"},
		{"Auth.WWWAuthenticate", headers.Auth.WWWAuthenticate(), "WWW-Authenticate"},
		{"Auth.ProxyAuthenticate", headers.Auth.ProxyAuthenticate(), "Proxy-Authenticate"},

		// Cache
		{"Cache.Age", headers.Cache.Age(), "Age"},
		{"Cache.CacheControl", headers.Cache.CacheControl(), "Cache-Control"},
		{"Cache.ClearSiteData", headers.Cache.ClearSiteData(), "Clear-Site-Data"},
		{"Cache.Expires", headers.Cache.Expires(), "Expires"},
		{"Cache.NoVarySearch", headers.Cache.NoVarySearch(), "No-Vary-Search"},
		{"Cache.Pragma", headers.Cache.Pragma(), "Pragma"},

		// Cond
		{"Cond.ETag", headers.Cond.ETag(), "ETag"},
		{"Cond.IfMatch", headers.Cond.IfMatch(), "If-Match"},
		{"Cond.IfNoneMatch", headers.Cond.IfNoneMatch(), "If-None-Match"},
		{"Cond.IfModifiedSince", headers.Cond.IfModifiedSince(), "If-Modified-Since"},
		{"Cond.IfUnmodifiedSince", headers.Cond.IfUnmodifiedSince(), "If-Unmodified-Since"},
		{"Cond.LastModified", headers.Cond.LastModified(), "Last-Modified"},
		{"Cond.Vary", headers.Cond.Vary(), "Vary"},

		// Conn
		{"Conn.Connection", headers.Conn.Connection(), "Connection"},
		{"Conn.KeepAlive", headers.Conn.KeepAlive(), "Keep-Alive"},

		// Negotiation
		{"Negotiation.Accept", headers.Negotiation.Accept(), "Accept"},
		{"Negotiation.AcceptEncoding", headers.Negotiation.AcceptEncoding(), "Accept-Encoding"},
		{"Negotiation.AcceptLanguage", headers.Negotiation.AcceptLanguage(), "Accept-Language"},
		{"Negotiation.AcceptPatch", headers.Negotiation.AcceptPatch(), "Accept-Patch"},
		{"Negotiation.AcceptPost", headers.Negotiation.AcceptPost(), "Accept-Post"},

		// Cookies
		{"Cookies.Cookie", headers.Cookies.Cookie(), "Cookie"},
		{"Cookies.SetCookie", headers.Cookies.SetCookie(), "Set-Cookie"},

		// CORS
		{"CORS.AllowCredentials", headers.CORS.AllowCredentials(), "Access-Control-Allow-Credentials"},
		{"CORS.AllowHeaders", headers.CORS.AllowHeaders(), "Access-Control-Allow-Headers"},
		{"CORS.AllowMethods", headers.CORS.AllowMethods(), "Access-Control-Allow-Methods"},
		{"CORS.AllowOrigin", headers.CORS.AllowOrigin(), "Access-Control-Allow-Origin"},
		{"CORS.ExposeHeaders", headers.CORS.ExposeHeaders(), "Access-Control-Expose-Headers"},
		{"CORS.MaxAge", headers.CORS.MaxAge(), "Access-Control-Max-Age"},
		{"CORS.RequestHeaders", headers.CORS.RequestHeaders(), "Access-Control-Request-Headers"},
		{"CORS.RequestMethod", headers.CORS.RequestMethod(), "Access-Control-Request-Method"},
		{"CORS.Origin", headers.CORS.Origin(), "Origin"},
		{"CORS.TimingAllowOrigin", headers.CORS.TimingAllowOrigin(), "Timing-Allow-Origin"},

		// Content
		{"Content.Disposition", headers.Content.Disposition(), "Content-Disposition"},
		{"Content.Encoding", headers.Content.Encoding(), "Content-Encoding"},
		{"Content.Language", headers.Content.Language(), "Content-Language"},
		{"Content.Length", headers.Content.Length(), "Content-Length"},
		{"Content.Location", headers.Content.Location(), "Content-Location"},
		{"Content.Type", headers.Content.Type(), "Content-Type"},
		{"Content.Range", headers.Content.Range(), "Content-Range"},

		// Ranges
		{"Ranges.AcceptRanges", headers.Ranges.AcceptRanges(), "Accept-Ranges"},
		{"Ranges.ContentRange", headers.Ranges.ContentRange(), "Content-Range"},
		{"Ranges.IfRange", headers.Ranges.IfRange(), "If-Range"},
		{"Ranges.Range", headers.Ranges.Range(), "Range"},

		// Redirect
		{"Redirect.Location", headers.Redirect.Location(), "Location"},
		{"Redirect.Refresh", headers.Redirect.Refresh(), "Refresh"},

		// Request
		{"Request.From", headers.Request.From(), "From"},
		{"Request.Host", headers.Request.Host(), "Host"},
		{"Request.Referer", headers.Request.Referer(), "Referer"},
		{"Request.ReferrerPolicy", headers.Request.ReferrerPolicy(), "Referrer-Policy"},
		{"Request.UserAgent", headers.Request.UserAgent(), "User-Agent"},

		// Response
		{"Response.Allow", headers.Response.Allow(), "Allow"},
		{"Response.Server", headers.Response.Server(), "Server"},

		// Security
		{"Security.CSP", headers.Security.CSP(), "Content-Security-Policy"},
		{"Security.CSPReportOnly", headers.Security.CSPReportOnly(), "Content-Security-Policy-Report-Only"},
		{"Security.COEP", headers.Security.COEP(), "Cross-Origin-Embedder-Policy"},
		{"Security.COOP", headers.Security.COOP(), "Cross-Origin-Opener-Policy"},
		{"Security.CORP", headers.Security.CORP(), "Cross-Origin-Resource-Policy"},
		{"Security.PermissionsPolicy", headers.Security.PermissionsPolicy(), "Permissions-Policy"},
		{"Security.HSTS", headers.Security.HSTS(), "Strict-Transport-Security"},
		{"Security.UpgradeInsecureRequests", headers.Security.UpgradeInsecureRequests(), "Upgrade-Insecure-Requests"},
		{"Security.XContentTypeOptions", headers.Security.XContentTypeOptions(), "X-Content-Type-Options"},
		{"Security.XFrameOptions", headers.Security.XFrameOptions(), "X-Frame-Options"},
		{"Security.XXSSProtection", headers.Security.XXSSProtection(), "X-XSS-Protection"},

		// WS
		{"WS.Accept", headers.WS.Accept(), "Sec-WebSocket-Accept"},
		{"WS.Extensions", headers.WS.Extensions(), "Sec-WebSocket-Extensions"},
		{"WS.Key", headers.WS.Key(), "Sec-WebSocket-Key"},
		{"WS.Protocol", headers.WS.Protocol(), "Sec-WebSocket-Protocol"},
		{"WS.Version", headers.WS.Version(), "Sec-WebSocket-Version"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("got %q, want %q", tt.got, tt.expected)
			}
		})
	}
}

func TestDirectHeaderConstants(t *testing.T) {
	// Sample some key constants to ensure they're exported correctly
	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"Authorization", headers.Authorization, "Authorization"},
		{"ContentType", headers.ContentType, "Content-Type"},
		{"CacheControl", headers.CacheControl, "Cache-Control"},
		{"UserAgent", headers.UserAgent, "User-Agent"},
		{"SetCookie", headers.SetCookie, "Set-Cookie"},
		{"AccessControlAllowOrigin", headers.AccessControlAllowOrigin, "Access-Control-Allow-Origin"},
		{"StrictTransportSecurity", headers.StrictTransportSecurity, "Strict-Transport-Security"},
		{"XForwardedFor", headers.XForwardedFor, "X-Forwarded-For"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("got %q, want %q", tt.got, tt.expected)
			}
		})
	}
}
