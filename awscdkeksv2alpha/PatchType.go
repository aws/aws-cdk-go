package awscdkeksv2alpha


// Values for `kubectl patch` --type argument.
// Experimental.
type PatchType string

const (
	// JSON Patch, RFC 6902.
	// Experimental.
	PatchType_JSON PatchType = "JSON"
	// JSON Merge patch.
	// Experimental.
	PatchType_MERGE PatchType = "MERGE"
	// Strategic merge patch.
	// Experimental.
	PatchType_STRATEGIC PatchType = "STRATEGIC"
)

