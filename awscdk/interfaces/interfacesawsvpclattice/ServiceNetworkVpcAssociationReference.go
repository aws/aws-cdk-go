package interfacesawsvpclattice


// A reference to a ServiceNetworkVpcAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNetworkVpcAssociationReference := &ServiceNetworkVpcAssociationReference{
//   	ServiceNetworkVpcAssociationArn: jsii.String("serviceNetworkVpcAssociationArn"),
//   }
//
type ServiceNetworkVpcAssociationReference struct {
	// The Arn of the ServiceNetworkVpcAssociation resource.
	ServiceNetworkVpcAssociationArn *string `field:"required" json:"serviceNetworkVpcAssociationArn" yaml:"serviceNetworkVpcAssociationArn"`
}

