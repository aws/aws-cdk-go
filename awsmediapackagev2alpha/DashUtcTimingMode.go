package awsmediapackagev2alpha


// The UTC timing mode for DASH.
// Experimental.
type DashUtcTimingMode string

const (
	// Option for HTTP_HEAD.
	// Experimental.
	DashUtcTimingMode_HTTP_HEAD DashUtcTimingMode = "HTTP_HEAD"
	// Option for HTTP_ISO.
	// Experimental.
	DashUtcTimingMode_HTTP_ISO DashUtcTimingMode = "HTTP_ISO"
	// Option for HTTP_XSDATE.
	// Experimental.
	DashUtcTimingMode_HTTP_XSDATE DashUtcTimingMode = "HTTP_XSDATE"
	// Option for UTC_DIRECT.
	// Experimental.
	DashUtcTimingMode_UTC_DIRECT DashUtcTimingMode = "UTC_DIRECT"
)

