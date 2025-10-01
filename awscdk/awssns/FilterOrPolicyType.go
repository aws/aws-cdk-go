package awssns


// The type of the MessageBody at a given key value pair.
type FilterOrPolicyType string

const (
	// The filter of the MessageBody.
	FilterOrPolicyType_FILTER FilterOrPolicyType = "FILTER"
	// A nested key of the MessageBody.
	FilterOrPolicyType_POLICY FilterOrPolicyType = "POLICY"
)

