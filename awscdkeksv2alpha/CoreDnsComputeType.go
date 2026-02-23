package awscdkeksv2alpha


// The type of compute resources to use for CoreDNS.
// Deprecated.
type CoreDnsComputeType string

const (
	// Deploy CoreDNS on EC2 instances.
	// Deprecated.
	CoreDnsComputeType_EC2 CoreDnsComputeType = "EC2"
	// Deploy CoreDNS on Fargate-managed instances.
	// Deprecated.
	CoreDnsComputeType_FARGATE CoreDnsComputeType = "FARGATE"
)

