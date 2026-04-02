package awsmediapackagev2alpha


// The input type will be an immutable field which will be used to define whether the channel will allow CMAF ingest or HLS ingest.
// Experimental.
type InputType string

const (
	// The DASH-IF CMAF Ingest specification (which defines CMAF segments with optional DASH manifests).
	// Experimental.
	InputType_CMAF InputType = "CMAF"
	// The HLS streaming specification (which defines M3U8 manifests and TS segments).
	// Experimental.
	InputType_HLS InputType = "HLS"
)

