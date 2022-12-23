package awsroute53


// The CAA tag.
type CaaTag string

const (
	// Explicity authorizes a single certificate authority to issue a certificate (any type) for the hostname.
	CaaTag_ISSUE CaaTag = "ISSUE"
	// Explicity authorizes a single certificate authority to issue a wildcard certificate (and only wildcard) for the hostname.
	CaaTag_ISSUEWILD CaaTag = "ISSUEWILD"
	// Specifies a URL to which a certificate authority may report policy violations.
	CaaTag_IODEF CaaTag = "IODEF"
)

