package awsnetworkmanager


// Properties for defining a `CfnCustomerGatewayAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCustomerGatewayAssociationProps := &cfnCustomerGatewayAssociationProps{
//   	customerGatewayArn: jsii.String("customerGatewayArn"),
//   	deviceId: jsii.String("deviceId"),
//   	globalNetworkId: jsii.String("globalNetworkId"),
//
//   	// the properties below are optional
//   	linkId: jsii.String("linkId"),
//   }
//
type CfnCustomerGatewayAssociationProps struct {
	// The Amazon Resource Name (ARN) of the customer gateway.
	CustomerGatewayArn *string `field:"required" json:"customerGatewayArn" yaml:"customerGatewayArn"`
	// The ID of the device.
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// The ID of the global network.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The ID of the link.
	LinkId *string `field:"optional" json:"linkId" yaml:"linkId"`
}

