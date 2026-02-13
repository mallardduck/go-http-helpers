# go-http-helpers

A collection of Go helper packages for working with HTTP requests and responses. Simplifies common HTTP operations with type-safe, convenient APIs.

## Features

- **headers**: Comprehensive HTTP header constants organized by context (CORS, Security, Auth, etc.)
- **query**: Type-safe URL query parameter extraction with automatic parsing and defaults

## Installation

```bash
go get github.com/mallardduck/go-http-helpers
```

## Packages

### headers

Provides HTTP header constants based on MDN Web Docs. Access headers either directly or through contextual groups for better discoverability.

#### Direct Access

```go
import "github.com/mallardduck/go-http-helpers/pkg/headers"

req.Header.Set(headers.ContentType, "application/json")
req.Header.Set(headers.Authorization, "Bearer token")
req.Header.Set(headers.CacheControl, "no-cache")
```

#### Grouped Access

```go
// Authentication
req.Header.Set(headers.Auth.Authorization(), "Bearer token")

// CORS
w.Header().Set(headers.CORS.AllowOrigin(), "*")
w.Header().Set(headers.CORS.AllowMethods(), "GET, POST, PUT")

// Security
w.Header().Set(headers.Security.CSP(), "default-src 'self'")
w.Header().Set(headers.Security.HSTS(), "max-age=31536000")

// Content
w.Header().Set(headers.Content.Type(), "text/html; charset=utf-8")
```

#### Available Groups

- **Auth**: Authentication headers
- **Cache**: Caching headers
- **Cond**: Conditional request headers
- **Conn**: Connection management
- **Negotiation**: Content negotiation
- **Cookies**: Cookie headers
- **CORS**: Cross-Origin Resource Sharing
- **Content**: Content-related headers
- **Ranges**: Range request headers
- **Redirect**: Redirect headers
- **Request**: Request context headers
- **Response**: Response context headers
- **Security**: Security-related headers
- **WS**: WebSocket headers

### query

Type-safe extraction and parsing of URL query parameters with automatic fallback to defaults.

#### Basic Usage

```go
import "github.com/mallardduck/go-http-helpers/pkg/query"

// URL: /products?page=2&limit=50&active=true&sort=price
page   := query.Int(r, "page", 1)           // 2
limit  := query.Int(r, "limit", 25)         // 50
active := query.Bool(r, "active", false)    // true
sort   := query.String(r, "sort", "name")   // "price"
offset := query.Int(r, "offset", 0)         // 0 (missing key)
```

#### Multiple Values

```go
// URL: /api?tag=go&tag=rust&tag=python
tags := query.Strings(r, "tag")  // []string{"go", "rust", "python"}

// URL: /filter?id=1&id=2&id=5
ids := query.Ints(r, "id", 0)    // []int{1, 2, 5}
```

#### Numeric Types

```go
count := query.Int(r, "count", 0)          // int
ratio := query.Float64(r, "ratio", 0.0)    // float64
id    := query.Int64(r, "id", 0)           // int64
```

#### Boolean Parsing

Flexibly parses common boolean representations (case-insensitive):

- **True**: "true", "1", "yes", "on", "y"
- **False**: "false", "0", "no", "off", "n"

```go
enabled := query.Bool(r, "enabled", false)
```

#### Generic Slice Parsing

```go
// URL: /api/items?id=1&id=2&id=invalid&id=5
ids := query.Slice(r, "id", 0, strconv.Atoi)
// Returns []int{1, 2, 0, 5} - invalid values use default

// Convenience helpers
ids := query.Ints(r, "id", 0)
values := query.Int64s(r, "val", -1)
prices := query.Float64s(r, "price", 0.0)
flags := query.Bools(r, "enabled", false)
```

#### Common Patterns

**Pagination:**
```go
page  := query.Int(r, "page", 1)
limit := query.Int(r, "limit", 25)
if page < 1 {
    page = 1
}
if limit < 1 || limit > 100 {
    limit = 25
}
```

**Search with Filters:**
```go
q          := query.String(r, "q", "")
category   := query.String(r, "category", "all")
tags       := query.Strings(r, "tag")
activeOnly := query.Bool(r, "active_only", false)
```

## Design Principles

- **Fail-safe**: Never panic on invalid input
- **Predictable**: Missing or invalid values always return the default
- **Concise**: One function call per parameter
- **Type-safe**: Return values match the requested type
- **Zero dependencies**: Uses only the standard library

## Requirements

- Go 1.24 or later

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.