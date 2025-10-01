package awsnetworkmanager


// A reference to a TransitGatewayRegistration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayRegistrationReference := &TransitGatewayRegistrationReference{
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	TransitGatewayArn: jsii.String("transitGatewayArn"),
//   }
//
type TransitGatewayRegistrationReference struct {
	// The GlobalNetworkId of the TransitGatewayRegistration resource.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The TransitGatewayArn of the TransitGatewayRegistration resource.
	TransitGatewayArn *string `field:"required" json:"transitGatewayArn" yaml:"transitGatewayArn"`
}

