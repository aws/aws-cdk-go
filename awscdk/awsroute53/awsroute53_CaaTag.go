package awsroute53


// The CAA tag.
// Experimental.
type CaaTag string

const (
	// Explicity authorizes a single certificate authority to issue a certificate (any type) for the hostname.
	// Experimental.
	CaaTag_ISSUE CaaTag = "ISSUE"
	// Explicity authorizes a single certificate authority to issue a wildcard certificate (and only wildcard) for the hostname.
	// Experimental.
	CaaTag_ISSUEWILD CaaTag = "ISSUEWILD"
	// Specifies a URL to which a certificate authority may report policy violations.
	// Experimental.
	CaaTag_IODEF CaaTag = "IODEF"
)

