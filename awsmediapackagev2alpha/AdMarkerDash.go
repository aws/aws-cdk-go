package awsmediapackagev2alpha


// Choose how ad markers are included in the packaged content.
// Experimental.
type AdMarkerDash string

const (
	// Option for BINARY - The SCTE-35 marker is expressed as a hex-string (Base64 string) rather than full XML.
	// Experimental.
	AdMarkerDash_BINARY AdMarkerDash = "BINARY"
	// Option for XML - The SCTE marker is expressed fully in XML.
	// Experimental.
	AdMarkerDash_XML AdMarkerDash = "XML"
)

