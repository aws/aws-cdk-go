package awsec2


// What action to apply to traffic matching the ACL.
// Experimental.
type Action string

const (
	// Allow the traffic.
	// Experimental.
	Action_ALLOW Action = "ALLOW"
	// Deny the traffic.
	// Experimental.
	Action_DENY Action = "DENY"
)

