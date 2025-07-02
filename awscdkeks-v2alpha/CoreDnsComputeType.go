package awscdkeks-v2alpha


// The type of compute resources to use for CoreDNS.
// Experimental.
type CoreDnsComputeType string

const (
	// Deploy CoreDNS on EC2 instances.
	// Experimental.
	CoreDnsComputeType_EC2 CoreDnsComputeType = "EC2"
	// Deploy CoreDNS on Fargate-managed instances.
	// Experimental.
	CoreDnsComputeType_FARGATE CoreDnsComputeType = "FARGATE"
)

