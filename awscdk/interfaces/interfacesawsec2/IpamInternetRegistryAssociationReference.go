package interfacesawsec2


// A reference to a IpamInternetRegistryAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipamInternetRegistryAssociationReference := &IpamInternetRegistryAssociationReference{
//   	IpamInternetRegistryAssociationArn: jsii.String("ipamInternetRegistryAssociationArn"),
//   }
//
type IpamInternetRegistryAssociationReference struct {
	// The IpamInternetRegistryAssociationArn of the IpamInternetRegistryAssociation resource.
	IpamInternetRegistryAssociationArn *string `field:"required" json:"ipamInternetRegistryAssociationArn" yaml:"ipamInternetRegistryAssociationArn"`
}

