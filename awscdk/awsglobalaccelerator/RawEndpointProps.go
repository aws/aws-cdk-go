package awsglobalaccelerator


// Properties for RawEndpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rawEndpointProps := &RawEndpointProps{
//   	EndpointId: jsii.String("endpointId"),
//
//   	// the properties below are optional
//   	PreserveClientIp: jsii.Boolean(false),
//   	Region: jsii.String("region"),
//   	Weight: jsii.Number(123),
//   }
//
type RawEndpointProps struct {
	// Identifier of the endpoint.
	//
	// Load balancer ARN, instance ID or EIP allocation ID.
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// Forward the client IP address.
	//
	// GlobalAccelerator will create Network Interfaces in your VPC in order
	// to preserve the client IP address.
	//
	// Only applies to Application Load Balancers and EC2 instances.
	//
	// Client IP address preservation is supported only in specific AWS Regions.
	// See the GlobalAccelerator Developer Guide for a list.
	// Default: true if possible and available.
	//
	PreserveClientIp *bool `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// The region where this endpoint is located.
	// Default: - Unknown what region this endpoint is located.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Endpoint weight across all endpoints in the group.
	//
	// Must be a value between 0 and 255.
	// Default: 128.
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

