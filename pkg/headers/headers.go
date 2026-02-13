// Package headers provides HTTP header constants organized by context.
// Headers can be accessed either directly or through contextual groups for discoverability.
package headers

// Individual header constants
const (
	// Authentication

	// Authorization contains the credentials to authenticate a user-agent with a server.
	Authorization = "Authorization"
	// ProxyAuthorization contains the credentials to authenticate a user agent with a proxy server.
	ProxyAuthorization = "Proxy-Authorization"
	// WWWAuthenticate defines the authentication method that should be used to access a resource.
	WWWAuthenticate = "WWW-Authenticate"
	// ProxyAuthenticate defines the authentication method that should be used to access a resource behind a proxy server.
	ProxyAuthenticate = "Proxy-Authenticate"

	// Caching

	// Age indicates the time, in seconds, that the object has been in a proxy cache.
	Age = "Age"
	// CacheControl specifies directives for caching mechanisms in both requests and responses.
	CacheControl = "Cache-Control"
	// ClearSiteData clears browsing data (e.g., cookies, storage, cache) associated with the requesting website.
	ClearSiteData = "Clear-Site-Data"
	// Expires indicates the date/time after which the response is considered stale.
	Expires = "Expires"
	// NoVarySearch specifies a set of rules that define how a URL's query parameters will affect cache matching.
	NoVarySearch = "No-Vary-Search"

	// Conditionals

	// ETag provides a unique string identifying the version of the resource.
	ETag = "ETag"
	// IfMatch makes the request conditional, and applies the method only if the stored resource matches one of the given ETags.
	IfMatch = "If-Match"
	// IfNoneMatch makes the request conditional, and applies the method only if the stored resource doesn't match any of the given ETags.
	IfNoneMatch = "If-None-Match"
	// IfModifiedSince makes the request conditional, and expects the resource to be transmitted only if it has been modified after the given date.
	IfModifiedSince = "If-Modified-Since"
	// IfUnmodifiedSince makes the request conditional, and expects the resource to be transmitted only if it has not been modified after the given date.
	IfUnmodifiedSince = "If-Unmodified-Since"
	// LastModified indicates the last modification date of the resource, used to compare several versions of the same resource.
	LastModified = "Last-Modified"
	// Vary determines how to match request headers to decide whether a cached response can be used rather than requesting a fresh one from the origin server.
	Vary = "Vary"

	// Connection Management

	// Connection controls whether the network connection stays open after the current transaction finishes.
	Connection = "Connection"
	// KeepAlive controls how long a persistent connection should stay open.
	KeepAlive = "Keep-Alive"

	// Content Negotiation

	// Accept informs the server about the types of data that can be sent back.
	Accept = "Accept"
	// AcceptEncoding specifies the encoding algorithm, usually a compression algorithm, that can be used on the resource sent back.
	AcceptEncoding = "Accept-Encoding"
	// AcceptLanguage informs the server about the human language the server is expected to send back.
	AcceptLanguage = "Accept-Language"
	// AcceptPatch is a request content negotiation response header that advertises which media type the server is able to understand in a PATCH request.
	AcceptPatch = "Accept-Patch"
	// AcceptPost is a request content negotiation response header that advertises which media type the server is able to understand in a POST request.
	AcceptPost = "Accept-Post"

	// Controls

	// Expect indicates expectations that need to be fulfilled by the server to properly handle the request.
	Expect = "Expect"
	// MaxForwards indicates the maximum number of hops the request can do before being reflected to the sender when using TRACE.
	MaxForwards = "Max-Forwards"

	// Cookies

	// Cookie contains stored HTTP cookies previously sent by the server with the Set-Cookie header.
	Cookie = "Cookie"
	// SetCookie sends cookies from the server to the user-agent.
	SetCookie = "Set-Cookie"

	// CORS

	// AccessControlAllowCredentials indicates whether the response to the request can be exposed when the credentials flag is true.
	AccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	// AccessControlAllowHeaders is used in response to a preflight request to indicate which HTTP headers can be used when making the actual request.
	AccessControlAllowHeaders = "Access-Control-Allow-Headers"
	// AccessControlAllowMethods specifies the methods allowed when accessing the resource in response to a preflight request.
	AccessControlAllowMethods = "Access-Control-Allow-Methods"
	// AccessControlAllowOrigin indicates whether the response can be shared.
	AccessControlAllowOrigin = "Access-Control-Allow-Origin"
	// AccessControlExposeHeaders indicates which headers can be exposed as part of the response by listing their names.
	AccessControlExposeHeaders = "Access-Control-Expose-Headers"
	// AccessControlMaxAge indicates how long the results of a preflight request can be cached.
	AccessControlMaxAge = "Access-Control-Max-Age"
	// AccessControlRequestHeaders is used when issuing a preflight request to let the server know which HTTP headers will be used when the actual request is made.
	AccessControlRequestHeaders = "Access-Control-Request-Headers"
	// AccessControlRequestMethod is used when issuing a preflight request to let the server know which HTTP method will be used when the actual request is made.
	AccessControlRequestMethod = "Access-Control-Request-Method"
	// Origin indicates where a fetch originates from.
	Origin = "Origin"
	// TimingAllowOrigin specifies origins that are allowed to see values of attributes retrieved via features of the Resource Timing API.
	TimingAllowOrigin = "Timing-Allow-Origin"

	// Downloads

	// ContentDisposition indicates if the resource transmitted should be displayed inline (default behavior without the header), or if it should be handled like a download.
	ContentDisposition = "Content-Disposition"

	// Integrity Digests

	// ContentDigest provides a digest of the stream of octets framed in an HTTP message (the message content) dependent on Content-Encoding and Content-Range.
	ContentDigest = "Content-Digest"
	// ReprDigest provides a digest of the selected representation of the target resource before transmission.
	ReprDigest = "Repr-Digest"
	// WantContentDigest states the wish for a Content-Digest header.
	WantContentDigest = "Want-Content-Digest"
	// WantReprDigest states the wish for a Repr-Digest header.
	WantReprDigest = "Want-Repr-Digest"

	// Message Body Information

	// ContentEncoding specifies the compression algorithm.
	ContentEncoding = "Content-Encoding"
	// ContentLanguage describes the human language(s) intended for the audience.
	ContentLanguage = "Content-Language"
	// ContentLength indicates the size of the resource, in decimal number of bytes.
	ContentLength = "Content-Length"
	// ContentLocation indicates an alternate location for the returned data.
	ContentLocation = "Content-Location"
	// ContentType indicates the media type of the resource.
	ContentType = "Content-Type"

	// Preferences

	// Prefer indicates preferences for specific server behaviors during request processing.
	Prefer = "Prefer"
	// PreferenceApplied informs the client which preferences specified in the Prefer header were applied by the server.
	PreferenceApplied = "Preference-Applied"

	// Proxies

	// Forwarded contains information from the client-facing side of proxy servers that is altered or lost when a proxy is involved in the path of the request.
	Forwarded = "Forwarded"
	// Via is added by proxies, both forward and reverse proxies, and can appear in the request headers and the response headers.
	Via = "Via"

	// Range Requests

	// AcceptRanges indicates if the server supports range requests, and if so in which unit the range can be expressed.
	AcceptRanges = "Accept-Ranges"
	// ContentRange indicates where in a full body message a partial message belongs.
	ContentRange = "Content-Range"
	// IfRange creates a conditional range request that is only fulfilled if the given etag or date matches the remote resource.
	IfRange = "If-Range"
	// Range indicates the part of a document that the server should return.
	Range = "Range"

	// Redirects

	// Location indicates the URL to redirect a page to.
	Location = "Location"
	// Refresh directs the browser to reload the page or redirect to another.
	Refresh = "Refresh"

	// Request Context

	// From contains an Internet email address for a human user who controls the requesting user agent.
	From = "From"
	// Host specifies the domain name of the server (for virtual hosting), and (optionally) the TCP port number on which the server is listening.
	Host = "Host"
	// Referer indicates the address of the previous web page from which a link to the currently requested page was followed.
	Referer = "Referer"
	// ReferrerPolicy governs which referrer information sent in the Referer header should be included with requests made.
	ReferrerPolicy = "Referrer-Policy"
	// UserAgent contains a characteristic string that allows the network protocol peers to identify the application type, operating system, software vendor or software version of the requesting software user agent.
	UserAgent = "User-Agent"

	// Response Context

	// Allow lists the set of HTTP request methods supported by a resource.
	Allow = "Allow"
	// Server contains information about the software used by the origin server to handle the request.
	Server = "Server"

	// Security

	// ContentSecurityPolicy controls resources the user agent is allowed to load for a given page.
	ContentSecurityPolicy = "Content-Security-Policy"
	// ContentSecurityPolicyReportOnly allows web developers to experiment with policies by monitoring, but not enforcing, their effects.
	ContentSecurityPolicyReportOnly = "Content-Security-Policy-Report-Only"
	// CrossOriginEmbedderPolicy allows a server to declare an embedder policy for a given document.
	CrossOriginEmbedderPolicy = "Cross-Origin-Embedder-Policy"
	// CrossOriginOpenerPolicy prevents other domains from opening/controlling a window.
	CrossOriginOpenerPolicy = "Cross-Origin-Opener-Policy"
	// CrossOriginResourcePolicy prevents other domains from reading the response of the resources to which this header is applied.
	CrossOriginResourcePolicy = "Cross-Origin-Resource-Policy"
	// PermissionsPolicy provides a mechanism to allow and deny the use of browser features in a website's own frame, and in iframes that it embeds.
	PermissionsPolicy = "Permissions-Policy"
	// ReportingEndpoints allows website owners to specify one or more endpoints used to receive errors such as CSP violation reports.
	ReportingEndpoints = "Reporting-Endpoints"
	// StrictTransportSecurity forces communication using HTTPS instead of HTTP.
	StrictTransportSecurity = "Strict-Transport-Security"
	// UpgradeInsecureRequests sends a signal to the server expressing the client's preference for an encrypted and authenticated response.
	UpgradeInsecureRequests = "Upgrade-Insecure-Requests"
	// XContentTypeOptions disables MIME sniffing and forces browser to use the type given in Content-Type.
	XContentTypeOptions = "X-Content-Type-Options"
	// XFrameOptions indicates whether a browser should be allowed to render a page in a frame, iframe, embed or object.
	XFrameOptions = "X-Frame-Options"
	// XPermittedCrossDomainPolicies overrides cross-domain policy files.
	XPermittedCrossDomainPolicies = "X-Permitted-Cross-Domain-Policies"
	// XPoweredBy may be set by hosting environments or other frameworks and contains information about them.
	XPoweredBy = "X-Powered-By"
	// XXSSProtection enables cross-site scripting filtering.
	XXSSProtection = "X-XSS-Protection"

	// Fetch Metadata

	// SecFetchDest indicates the request's destination.
	SecFetchDest = "Sec-Fetch-Dest"
	// SecFetchMode indicates the request's mode to a server.
	SecFetchMode = "Sec-Fetch-Mode"
	// SecFetchSite indicates the relationship between a request initiator's origin and its target's origin.
	SecFetchSite = "Sec-Fetch-Site"
	// SecFetchUser indicates whether or not a navigation request was triggered by user activation.
	SecFetchUser = "Sec-Fetch-User"
	// SecPurpose indicates the purpose of the request, when the purpose is something other than immediate use by the user-agent.
	SecPurpose = "Sec-Purpose"

	// Fetch Storage Access

	// SecFetchStorageAccess indicates the "storage access status" for the current fetch context.
	SecFetchStorageAccess = "Sec-Fetch-Storage-Access"
	// ActivateStorageAccess is used in response to Sec-Fetch-Storage-Access to indicate that the browser can activate an existing permission.
	ActivateStorageAccess = "Activate-Storage-Access"

	// Server-Sent Events

	// ReportTo is used to specify server endpoints where the browser should send warning and error reports.
	ReportTo = "Report-To"

	// Transfer Coding

	// TE specifies the transfer encodings the user agent is willing to accept.
	TE = "TE"
	// Trailer allows the sender to include additional fields at the end of chunked message.
	Trailer = "Trailer"
	// TransferEncoding specifies the form of encoding used to safely transfer the resource to the user.
	TransferEncoding = "Transfer-Encoding"

	// WebSockets

	// SecWebSocketAccept indicates that the server is willing to upgrade to a WebSocket connection.
	SecWebSocketAccept = "Sec-WebSocket-Accept"
	// SecWebSocketExtensions in requests, indicates the WebSocket extensions supported by the client; in responses, indicates the extension selected by the server.
	SecWebSocketExtensions = "Sec-WebSocket-Extensions"
	// SecWebSocketKey contains a key that verifies that the client explicitly intends to open a WebSocket.
	SecWebSocketKey = "Sec-WebSocket-Key"
	// SecWebSocketProtocol in requests, indicates the sub-protocols supported by the client; in responses, indicates the sub-protocol selected by the server.
	SecWebSocketProtocol = "Sec-WebSocket-Protocol"
	// SecWebSocketVersion indicates the version of the WebSocket protocol used by the client.
	SecWebSocketVersion = "Sec-WebSocket-Version"

	// Other

	// AltSvc is used to list alternate ways to reach this service.
	AltSvc = "Alt-Svc"
	// AltUsed is used to identify the alternative service in use.
	AltUsed = "Alt-Used"
	// Date contains the date and time at which the message was originated.
	Date = "Date"
	// Link provides a means for serializing one or more links in HTTP headers.
	Link = "Link"
	// RetryAfter indicates how long the user agent should wait before making a follow-up request.
	RetryAfter = "Retry-After"
	// ServerTiming communicates one or more metrics and descriptions for the given request-response cycle.
	ServerTiming = "Server-Timing"
	// ServiceWorker is included in fetches for a service worker's script resource.
	ServiceWorker = "Service-Worker"
	// ServiceWorkerAllowed is used to remove the path restriction by including this header in the response of the Service Worker script.
	ServiceWorkerAllowed = "Service-Worker-Allowed"
	// ServiceWorkerNavigationPreload is sent in preemptive request to fetch a resource during service worker boot.
	ServiceWorkerNavigationPreload = "Service-Worker-Navigation-Preload"
	// SourceMap links to a source map so that debuggers can step through original source code instead of generated or transformed code.
	SourceMap = "SourceMap"
	// Upgrade can be used to upgrade an already established client/server connection to a different protocol.
	Upgrade = "Upgrade"
	// Priority provides a hint from about the priority of a particular resource request on a particular connection.
	Priority = "Priority"

	// Client Hints

	// AcceptCH allows servers to advertise support for Client Hints.
	AcceptCH = "Accept-CH"
	// CriticalCH is used along with Accept-CH to specify that accepted client hints are also critical client hints.
	CriticalCH = "Critical-CH"
	// SecCHUA indicates the user agent's branding and version.
	SecCHUA = "Sec-CH-UA"
	// SecCHUAArch indicates the user agent's underlying platform architecture.
	SecCHUAArch = "Sec-CH-UA-Arch"
	// SecCHUABitness indicates the user agent's underlying CPU architecture bitness.
	SecCHUABitness = "Sec-CH-UA-Bitness"
	// SecCHUAFormFactors indicates the user agent's form-factors.
	SecCHUAFormFactors = "Sec-CH-UA-Form-Factors"
	// SecCHUAFullVersion indicates the user agent's full version string.
	SecCHUAFullVersion = "Sec-CH-UA-Full-Version"
	// SecCHUAFullVersionList indicates the full version for each brand in the user agent's brand list.
	SecCHUAFullVersionList = "Sec-CH-UA-Full-Version-List"
	// SecCHUAMobile indicates whether the user agent is running on a mobile device or prefers a "mobile" user experience.
	SecCHUAMobile = "Sec-CH-UA-Mobile"
	// SecCHUAModel indicates the user agent's device model.
	SecCHUAModel = "Sec-CH-UA-Model"
	// SecCHUAPlatform indicates the user agent's underlying operation system/platform.
	SecCHUAPlatform = "Sec-CH-UA-Platform"
	// SecCHUAPlatformVersion indicates the user agent's underlying operation system version.
	SecCHUAPlatformVersion = "Sec-CH-UA-Platform-Version"
	// SecCHUAWoW64 indicates whether or not the user agent binary is running in 32-bit mode on 64-bit Windows.
	SecCHUAWoW64 = "Sec-CH-UA-WoW64"
	// SecCHPrefersColorScheme indicates the user's preference of dark or light color scheme.
	SecCHPrefersColorScheme = "Sec-CH-Prefers-Color-Scheme"
	// SecCHPrefersReducedMotion indicates the user's preference to see fewer animations and content layout shifts.
	SecCHPrefersReducedMotion = "Sec-CH-Prefers-Reduced-Motion"
	// SecCHPrefersReducedTransparency indicates the user agent's preference for reduced transparency.
	SecCHPrefersReducedTransparency = "Sec-CH-Prefers-Reduced-Transparency"
	// SecCHDeviceMemory indicates the approximate amount of available client RAM memory.
	SecCHDeviceMemory = "Sec-CH-Device-Memory"
	// SecCHDPR indicates the client device pixel ratio.
	SecCHDPR = "Sec-CH-DPR"
	// SecCHViewportHeight indicates the client's layout viewport height in CSS pixels.
	SecCHViewportHeight = "Sec-CH-Viewport-Height"
	// SecCHViewportWidth indicates the client's layout viewport width in CSS pixels.
	SecCHViewportWidth = "Sec-CH-Viewport-Width"
	// Downlink indicates the approximate bandwidth of the client's connection to the server, in Mbps.
	Downlink = "Downlink"
	// ECT indicates the effective connection type that best matches the connection's latency and bandwidth.
	ECT = "ECT"
	// RTT indicates the application layer round trip time (RTT) in milliseconds.
	RTT = "RTT"
	// SaveData indicates the user agent's preference for reduced data usage.
	SaveData = "Save-Data"

	// Compression Dictionary Transport

	// AvailableDictionary indicates the best dictionary a browser has available for the server to use for compression.
	AvailableDictionary = "Available-Dictionary"
	// DictionaryID is used when a browser already has a dictionary available for a resource.
	DictionaryID = "Dictionary-ID"
	// UseAsDictionary lists the matching criteria that the dictionary can be used for in future requests.
	UseAsDictionary = "Use-As-Dictionary"

	// Privacy

	// DNT indicates the user's tracking preference (Do Not Track).
	DNT = "DNT"
	// Tk indicates the tracking status that applied to the corresponding request.
	Tk = "Tk"
	// SecGPC indicates whether the user consents to a website or service selling or sharing their personal information with third parties.
	SecGPC = "Sec-GPC"

	// Non-standard but common

	// XForwardedFor identifies the originating IP addresses of a client connecting to a web server through an HTTP proxy or a load balancer.
	XForwardedFor = "X-Forwarded-For"
	// XForwardedHost identifies the original host requested that a client used to connect to your proxy or load balancer.
	XForwardedHost = "X-Forwarded-Host"
	// XForwardedProto identifies the protocol (HTTP or HTTPS) that a client used to connect to your proxy or load balancer.
	XForwardedProto = "X-Forwarded-Proto"
	// XDNSPrefetchControl controls DNS prefetching.
	XDNSPrefetchControl = "X-DNS-Prefetch-Control"
	// XRobotsTag indicates how a web page is to be indexed within public search engine results.
	XRobotsTag = "X-Robots-Tag"

	// Deprecated

	// Pragma is an implementation-specific header that may have various effects anywhere along the request-response chain.
	Pragma = "Pragma"
	// Warning provides general warning information about possible problems.
	Warning = "Warning"
)

// Auth provides authentication-related headers
type authHeaders struct{}

var Auth = authHeaders{}

func (authHeaders) Authorization() string      { return Authorization }
func (authHeaders) ProxyAuthorization() string { return ProxyAuthorization }
func (authHeaders) WWWAuthenticate() string    { return WWWAuthenticate }
func (authHeaders) ProxyAuthenticate() string  { return ProxyAuthenticate }

// Cache provides caching-related headers
type cacheHeaders struct{}

var Cache = cacheHeaders{}

func (cacheHeaders) Age() string           { return Age }
func (cacheHeaders) CacheControl() string  { return CacheControl }
func (cacheHeaders) ClearSiteData() string { return ClearSiteData }
func (cacheHeaders) Expires() string       { return Expires }
func (cacheHeaders) NoVarySearch() string  { return NoVarySearch }
func (cacheHeaders) Pragma() string        { return Pragma }

// Cond provides conditional request headers
type condHeaders struct{}

var Cond = condHeaders{}

func (condHeaders) ETag() string              { return ETag }
func (condHeaders) IfMatch() string           { return IfMatch }
func (condHeaders) IfNoneMatch() string       { return IfNoneMatch }
func (condHeaders) IfModifiedSince() string   { return IfModifiedSince }
func (condHeaders) IfUnmodifiedSince() string { return IfUnmodifiedSince }
func (condHeaders) LastModified() string      { return LastModified }
func (condHeaders) Vary() string              { return Vary }

// Conn provides connection management headers
type connHeaders struct{}

var Conn = connHeaders{}

func (connHeaders) Connection() string { return Connection }
func (connHeaders) KeepAlive() string  { return KeepAlive }

// Negotiation provides content negotiation headers
type negotiationHeaders struct{}

var Negotiation = negotiationHeaders{}

func (negotiationHeaders) Accept() string         { return Accept }
func (negotiationHeaders) AcceptEncoding() string { return AcceptEncoding }
func (negotiationHeaders) AcceptLanguage() string { return AcceptLanguage }
func (negotiationHeaders) AcceptPatch() string    { return AcceptPatch }
func (negotiationHeaders) AcceptPost() string     { return AcceptPost }

// Cookies provides cookie-related headers
type cookieHeaders struct{}

var Cookies = cookieHeaders{}

func (cookieHeaders) Cookie() string    { return Cookie }
func (cookieHeaders) SetCookie() string { return SetCookie }

// CORS provides Cross-Origin Resource Sharing headers
type corsHeaders struct{}

var CORS = corsHeaders{}

func (corsHeaders) AllowCredentials() string  { return AccessControlAllowCredentials }
func (corsHeaders) AllowHeaders() string      { return AccessControlAllowHeaders }
func (corsHeaders) AllowMethods() string      { return AccessControlAllowMethods }
func (corsHeaders) AllowOrigin() string       { return AccessControlAllowOrigin }
func (corsHeaders) ExposeHeaders() string     { return AccessControlExposeHeaders }
func (corsHeaders) MaxAge() string            { return AccessControlMaxAge }
func (corsHeaders) RequestHeaders() string    { return AccessControlRequestHeaders }
func (corsHeaders) RequestMethod() string     { return AccessControlRequestMethod }
func (corsHeaders) Origin() string            { return Origin }
func (corsHeaders) TimingAllowOrigin() string { return TimingAllowOrigin }

// Content provides content-related headers
type contentHeaders struct{}

var Content = contentHeaders{}

func (contentHeaders) Disposition() string { return ContentDisposition }
func (contentHeaders) Encoding() string    { return ContentEncoding }
func (contentHeaders) Language() string    { return ContentLanguage }
func (contentHeaders) Length() string      { return ContentLength }
func (contentHeaders) Location() string    { return ContentLocation }
func (contentHeaders) Type() string        { return ContentType }
func (contentHeaders) Range() string       { return ContentRange }

// Range provides range request headers
type rangeHeaders struct{}

var Ranges = rangeHeaders{}

func (rangeHeaders) AcceptRanges() string { return AcceptRanges }
func (rangeHeaders) ContentRange() string { return ContentRange }
func (rangeHeaders) IfRange() string      { return IfRange }
func (rangeHeaders) Range() string        { return Range }

// Redirect provides redirect-related headers
type redirectHeaders struct{}

var Redirect = redirectHeaders{}

func (redirectHeaders) Location() string { return Location }
func (redirectHeaders) Refresh() string  { return Refresh }

// Request provides request context headers
type requestHeaders struct{}

var Request = requestHeaders{}

func (requestHeaders) From() string           { return From }
func (requestHeaders) Host() string           { return Host }
func (requestHeaders) Referer() string        { return Referer }
func (requestHeaders) ReferrerPolicy() string { return ReferrerPolicy }
func (requestHeaders) UserAgent() string      { return UserAgent }

// Response provides response context headers
type responseHeaders struct{}

var Response = responseHeaders{}

func (responseHeaders) Allow() string  { return Allow }
func (responseHeaders) Server() string { return Server }

// Security provides security-related headers
type securityHeaders struct{}

var Security = securityHeaders{}

func (securityHeaders) CSP() string                     { return ContentSecurityPolicy }
func (securityHeaders) CSPReportOnly() string           { return ContentSecurityPolicyReportOnly }
func (securityHeaders) COEP() string                    { return CrossOriginEmbedderPolicy }
func (securityHeaders) COOP() string                    { return CrossOriginOpenerPolicy }
func (securityHeaders) CORP() string                    { return CrossOriginResourcePolicy }
func (securityHeaders) PermissionsPolicy() string       { return PermissionsPolicy }
func (securityHeaders) HSTS() string                    { return StrictTransportSecurity }
func (securityHeaders) UpgradeInsecureRequests() string { return UpgradeInsecureRequests }
func (securityHeaders) XContentTypeOptions() string     { return XContentTypeOptions }
func (securityHeaders) XFrameOptions() string           { return XFrameOptions }
func (securityHeaders) XXSSProtection() string          { return XXSSProtection }

// WS provides WebSocket headers
type wsHeaders struct{}

var WS = wsHeaders{}

func (wsHeaders) Accept() string     { return SecWebSocketAccept }
func (wsHeaders) Extensions() string { return SecWebSocketExtensions }
func (wsHeaders) Key() string        { return SecWebSocketKey }
func (wsHeaders) Protocol() string   { return SecWebSocketProtocol }
func (wsHeaders) Version() string    { return SecWebSocketVersion }
