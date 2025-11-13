package interfacesawsvpclattice


// A reference to a ServiceNetworkResourceAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNetworkResourceAssociationReference := &ServiceNetworkResourceAssociationReference{
//   	ServiceNetworkResourceAssociationArn: jsii.String("serviceNetworkResourceAssociationArn"),
//   }
//
type ServiceNetworkResourceAssociationReference struct {
	// The Arn of the ServiceNetworkResourceAssociation resource.
	ServiceNetworkResourceAssociationArn *string `field:"required" json:"serviceNetworkResourceAssociationArn" yaml:"serviceNetworkResourceAssociationArn"`
}

