package awsappmesh


// A reference to a VirtualGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualGatewayReference := &VirtualGatewayReference{
//   	VirtualGatewayArn: jsii.String("virtualGatewayArn"),
//   	VirtualGatewayId: jsii.String("virtualGatewayId"),
//   }
//
type VirtualGatewayReference struct {
	// The ARN of the VirtualGateway resource.
	VirtualGatewayArn *string `field:"required" json:"virtualGatewayArn" yaml:"virtualGatewayArn"`
	// The Id of the VirtualGateway resource.
	VirtualGatewayId *string `field:"required" json:"virtualGatewayId" yaml:"virtualGatewayId"`
}

