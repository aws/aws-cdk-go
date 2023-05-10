package awsec2


// Direction of traffic the AclEntry applies to.
// Experimental.
type TrafficDirection string

const (
	// Traffic leaving the subnet.
	// Experimental.
	TrafficDirection_EGRESS TrafficDirection = "EGRESS"
	// Traffic entering the subnet.
	// Experimental.
	TrafficDirection_INGRESS TrafficDirection = "INGRESS"
)

