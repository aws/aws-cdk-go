package awsec2alpha


// Limits which service in AWS that the pool can be used in.
// Experimental.
type AwsServiceName string

const (
	// Allows users to use space for Elastic IP addresses and VPCs.
	// Experimental.
	AwsServiceName_EC2 AwsServiceName = "EC2"
)

