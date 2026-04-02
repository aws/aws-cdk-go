package awsmediapackagev2alpha


// Controls whether SCTE-35 messages are included in segment files.
// Experimental.
type ScteInSegments string

const (
	// SCTE-35 messages are not included in segments.
	// Experimental.
	ScteInSegments_NONE ScteInSegments = "NONE"
	// SCTE-35 messages are embedded in segment data.
	//
	// For DASH manifests, when set to ALL, an InbandEventStream tag signals
	// that SCTE messages are present in segments.
	// Experimental.
	ScteInSegments_ALL ScteInSegments = "ALL"
)

