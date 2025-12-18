package interfacesawsnetworkmanager


// A reference to a CoreNetworkPrefixListAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   coreNetworkPrefixListAssociationReference := &CoreNetworkPrefixListAssociationReference{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	PrefixListArn: jsii.String("prefixListArn"),
//   }
//
type CoreNetworkPrefixListAssociationReference struct {
	// The CoreNetworkId of the CoreNetworkPrefixListAssociation resource.
	CoreNetworkId *string `field:"required" json:"coreNetworkId" yaml:"coreNetworkId"`
	// The PrefixListArn of the CoreNetworkPrefixListAssociation resource.
	PrefixListArn *string `field:"required" json:"prefixListArn" yaml:"prefixListArn"`
}

