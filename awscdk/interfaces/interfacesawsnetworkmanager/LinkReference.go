package interfacesawsnetworkmanager


// A reference to a Link resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   linkReference := &LinkReference{
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	LinkArn: jsii.String("linkArn"),
//   	LinkId: jsii.String("linkId"),
//   }
//
type LinkReference struct {
	// The GlobalNetworkId of the Link resource.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// The ARN of the Link resource.
	LinkArn *string `field:"required" json:"linkArn" yaml:"linkArn"`
	// The LinkId of the Link resource.
	LinkId *string `field:"required" json:"linkId" yaml:"linkId"`
}

