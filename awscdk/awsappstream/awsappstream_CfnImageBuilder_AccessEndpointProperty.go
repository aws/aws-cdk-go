package awsappstream


// Describes an interface VPC endpoint (interface endpoint) that lets you create a private connection between the virtual private cloud (VPC) that you specify and AppStream 2.0. When you specify an interface endpoint for a stack, users of the stack can connect to AppStream 2.0 only through that endpoint. When you specify an interface endpoint for an image builder, administrators can connect to the image builder only through that endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessEndpointProperty := &accessEndpointProperty{
//   	endpointType: jsii.String("endpointType"),
//   	vpceId: jsii.String("vpceId"),
//   }
//
type CfnImageBuilder_AccessEndpointProperty struct {
	// The type of interface endpoint.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// The identifier (ID) of the VPC in which the interface endpoint is used.
	VpceId *string `field:"required" json:"vpceId" yaml:"vpceId"`
}

