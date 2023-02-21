package awseks


// Values for `kubectl patch` --type argument.
type PatchType string

const (
	// JSON Patch, RFC 6902.
	PatchType_JSON PatchType = "JSON"
	// JSON Merge patch.
	PatchType_MERGE PatchType = "MERGE"
	// Strategic merge patch.
	PatchType_STRATEGIC PatchType = "STRATEGIC"
)

