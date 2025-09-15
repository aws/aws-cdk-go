package awsopensearchserverless


// A reference to a VpcEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcEndpointReference := &VpcEndpointReference{
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   }
//
type VpcEndpointReference struct {
	// The Id of the VpcEndpoint resource.
	VpcEndpointId *string `field:"required" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

