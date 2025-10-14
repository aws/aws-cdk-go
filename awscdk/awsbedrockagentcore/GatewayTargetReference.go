package awsbedrockagentcore


// A reference to a GatewayTarget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayTargetReference := &GatewayTargetReference{
//   	GatewayIdentifier: jsii.String("gatewayIdentifier"),
//   	TargetId: jsii.String("targetId"),
//   }
//
type GatewayTargetReference struct {
	// The GatewayIdentifier of the GatewayTarget resource.
	GatewayIdentifier *string `field:"required" json:"gatewayIdentifier" yaml:"gatewayIdentifier"`
	// The TargetId of the GatewayTarget resource.
	TargetId *string `field:"required" json:"targetId" yaml:"targetId"`
}

