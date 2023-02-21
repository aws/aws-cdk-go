package awseks


// The type of compute resources to use for CoreDNS.
type CoreDnsComputeType string

const (
	// Deploy CoreDNS on EC2 instances.
	CoreDnsComputeType_EC2 CoreDnsComputeType = "EC2"
	// Deploy CoreDNS on Fargate-managed instances.
	CoreDnsComputeType_FARGATE CoreDnsComputeType = "FARGATE"
)

