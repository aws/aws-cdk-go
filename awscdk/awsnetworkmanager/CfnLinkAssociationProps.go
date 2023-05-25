package awsnetworkmanager


// Properties for defining a `CfnLinkAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLinkAssociationProps := &CfnLinkAssociationProps{
//   	DeviceId: jsii.String("deviceId"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	LinkId: jsii.String("linkId"),
//   }
//
type CfnLinkAssociationProps struct {
	// The device ID for the link association.
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// The ID of the global network.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The ID of the link.
	LinkId *string `field:"required" json:"linkId" yaml:"linkId"`
}

