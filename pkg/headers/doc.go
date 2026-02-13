// Package headers provides HTTP header constants organized by context.
//
// This package offers a comprehensive set of HTTP header constants based on
// the Mozilla Developer Network (MDN) Web Docs. Headers can be accessed either
// directly as constants or through contextual groups for better discoverability.
//
// # Direct Access
//
// All header constants are available as package-level constants:
//
//	req.Header.Set(headers.ContentType, "application/json")
//	req.Header.Set(headers.Authorization, "Bearer token")
//	req.Header.Set(headers.CacheControl, "no-cache")
//
// # Grouped Access
//
// Headers are also organized into contextual groups for improved discoverability
// and code organization:
//
//	// Authentication headers
//	req.Header.Set(headers.Auth.Authorization(), "Bearer token")
//	req.Header.Set(headers.Auth.WWWAuthenticate(), "Basic realm=\"api\"")
//
//	// CORS headers
//	w.Header().Set(headers.CORS.AllowOrigin(), "*")
//	w.Header().Set(headers.CORS.AllowMethods(), "GET, POST, PUT")
//
//	// Security headers
//	w.Header().Set(headers.Security.CSP(), "default-src 'self'")
//	w.Header().Set(headers.Security.HSTS(), "max-age=31536000")
//
//	// Content headers
//	w.Header().Set(headers.Content.Type(), "text/html; charset=utf-8")
//	w.Header().Set(headers.Content.Length(), "1234")
//
// # Available Groups
//
// The following header groups are available:
//
//   - Auth: Authentication headers (Authorization, WWW-Authenticate, etc.)
//   - Cache: Caching headers (Cache-Control, Expires, Age, etc.)
//   - Cond: Conditional request headers (ETag, If-Match, If-None-Match, etc.)
//   - Conn: Connection management headers (Connection, Keep-Alive)
//   - Negotiation: Content negotiation headers (Accept, Accept-Encoding, etc.)
//   - Cookies: Cookie headers (Cookie, Set-Cookie)
//   - CORS: Cross-Origin Resource Sharing headers
//   - Content: Content-related headers (Content-Type, Content-Length, etc.)
//   - Ranges: Range request headers (Range, Content-Range, etc.)
//   - Redirect: Redirect headers (Location, Refresh)
//   - Request: Request context headers (Host, User-Agent, Referer, etc.)
//   - Response: Response context headers (Allow, Server)
//   - Security: Security-related headers (CSP, HSTS, XFO, etc.)
//   - WS: WebSocket headers (Sec-WebSocket-Key, Sec-WebSocket-Accept, etc.)
//
// # Header Values
//
// All header constant values match the official HTTP header specifications
// as documented on MDN. While Go's net/http package internally canonicalizes
// header keys, these constants use the standard documented format.
//
// # Standards Compliance
//
// The header constants and descriptions are based on the MDN Web Docs:
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
package headers
