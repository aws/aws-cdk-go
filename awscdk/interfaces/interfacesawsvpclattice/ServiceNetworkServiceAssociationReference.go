package interfacesawsvpclattice


// A reference to a ServiceNetworkServiceAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceNetworkServiceAssociationReference := &ServiceNetworkServiceAssociationReference{
//   	ServiceNetworkServiceAssociationArn: jsii.String("serviceNetworkServiceAssociationArn"),
//   }
//
type ServiceNetworkServiceAssociationReference struct {
	// The Arn of the ServiceNetworkServiceAssociation resource.
	ServiceNetworkServiceAssociationArn *string `field:"required" json:"serviceNetworkServiceAssociationArn" yaml:"serviceNetworkServiceAssociationArn"`
}

