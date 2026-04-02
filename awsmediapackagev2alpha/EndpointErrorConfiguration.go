package awsmediapackagev2alpha


// Endpoint error configuration options.
// Experimental.
type EndpointErrorConfiguration string

const (
	// The manifest stalled and there are no new segments or parts.
	// Experimental.
	EndpointErrorConfiguration_STALE_MANIFEST EndpointErrorConfiguration = "STALE_MANIFEST"
	// There is a gap in the manifest.
	// Experimental.
	EndpointErrorConfiguration_INCOMPLETE_MANIFEST EndpointErrorConfiguration = "INCOMPLETE_MANIFEST"
	// Key rotation is enabled but we're unable to fetch the key for the current key period.
	// Experimental.
	EndpointErrorConfiguration_MISSING_DRM_KEY EndpointErrorConfiguration = "MISSING_DRM_KEY"
	// The segments which contain slate content are considered to be missing content.
	// Experimental.
	EndpointErrorConfiguration_SLATE_INPUT EndpointErrorConfiguration = "SLATE_INPUT"
)

