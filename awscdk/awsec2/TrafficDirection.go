package awsec2


// Direction of traffic the AclEntry applies to.
type TrafficDirection string

const (
	// Traffic leaving the subnet.
	TrafficDirection_EGRESS TrafficDirection = "EGRESS"
	// Traffic entering the subnet.
	TrafficDirection_INGRESS TrafficDirection = "INGRESS"
)

