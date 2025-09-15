package awsec2


// A reference to a VPCEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPCEndpointReference := &VPCEndpointReference{
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   }
//
type VPCEndpointReference struct {
	// The Id of the VPCEndpoint resource.
	VpcEndpointId *string `field:"required" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

