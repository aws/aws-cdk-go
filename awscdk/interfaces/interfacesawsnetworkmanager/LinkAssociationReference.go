package interfacesawsnetworkmanager


// A reference to a LinkAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linkAssociationReference := &LinkAssociationReference{
//   	DeviceId: jsii.String("deviceId"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	LinkId: jsii.String("linkId"),
//   }
//
type LinkAssociationReference struct {
	// The DeviceId of the LinkAssociation resource.
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// The GlobalNetworkId of the LinkAssociation resource.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The LinkId of the LinkAssociation resource.
	LinkId *string `field:"required" json:"linkId" yaml:"linkId"`
}

