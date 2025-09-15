package awsec2


// A reference to a SubnetNetworkAclAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetNetworkAclAssociationReference := &SubnetNetworkAclAssociationReference{
//   	AssociationId: jsii.String("associationId"),
//   }
//
type SubnetNetworkAclAssociationReference struct {
	// The AssociationId of the SubnetNetworkAclAssociation resource.
	AssociationId *string `field:"required" json:"associationId" yaml:"associationId"`
}

