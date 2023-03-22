package awsnetworkmanager


// Properties for defining a `CfnTransitGatewayRegistration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTransitGatewayRegistrationProps := &cfnTransitGatewayRegistrationProps{
//   	globalNetworkId: jsii.String("globalNetworkId"),
//   	transitGatewayArn: jsii.String("transitGatewayArn"),
//   }
//
type CfnTransitGatewayRegistrationProps struct {
	// The ID of the global network.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The Amazon Resource Name (ARN) of the transit gateway.
	TransitGatewayArn *string `field:"required" json:"transitGatewayArn" yaml:"transitGatewayArn"`
}

