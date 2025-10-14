package awsbedrockagentcore


// A reference to a Gateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayReference := &GatewayReference{
//   	GatewayArn: jsii.String("gatewayArn"),
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   }
//
type GatewayReference struct {
	// The ARN of the Gateway resource.
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// The GatewayIdentifier of the Gateway resource.
	GatewayIdentifier *string `field:"required" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
}

