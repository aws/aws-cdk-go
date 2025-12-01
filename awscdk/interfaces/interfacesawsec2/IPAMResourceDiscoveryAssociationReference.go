package interfacesawsec2


// A reference to a IPAMResourceDiscoveryAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPAMResourceDiscoveryAssociationReference := map[string]*string{
//   	"ipamResourceDiscoveryAssociationArn": jsii.String("ipamResourceDiscoveryAssociationArn"),
//   	"ipamResourceDiscoveryAssociationId": jsii.String("ipamResourceDiscoveryAssociationId"),
//   }
//
type IPAMResourceDiscoveryAssociationReference struct {
	// The ARN of the IPAMResourceDiscoveryAssociation resource.
	IpamResourceDiscoveryAssociationArn *string `field:"required" json:"ipamResourceDiscoveryAssociationArn" yaml:"ipamResourceDiscoveryAssociationArn"`
	// The IpamResourceDiscoveryAssociationId of the IPAMResourceDiscoveryAssociation resource.
	IpamResourceDiscoveryAssociationId *string `field:"required" json:"ipamResourceDiscoveryAssociationId" yaml:"ipamResourceDiscoveryAssociationId"`
}

