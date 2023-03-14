// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// Platform capabilities.
// Experimental.
type PlatformCapabilities string

const (
	// Specifies EC2 environment.
	// Experimental.
	PlatformCapabilities_EC2 PlatformCapabilities = "EC2"
	// Specifies Fargate environment.
	// Experimental.
	PlatformCapabilities_FARGATE PlatformCapabilities = "FARGATE"
)

