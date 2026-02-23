package awscdkeksv2alpha


// Values for `kubectl patch` --type argument.
// Deprecated.
type PatchType string

const (
	// JSON Patch, RFC 6902.
	// Deprecated.
	PatchType_JSON PatchType = "JSON"
	// JSON Merge patch.
	// Deprecated.
	PatchType_MERGE PatchType = "MERGE"
	// Strategic merge patch.
	// Deprecated.
	PatchType_STRATEGIC PatchType = "STRATEGIC"
)

