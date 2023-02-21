package awsec2


// What action to apply to traffic matching the ACL.
type Action string

const (
	// Allow the traffic.
	Action_ALLOW Action = "ALLOW"
	// Deny the traffic.
	Action_DENY Action = "DENY"
)

