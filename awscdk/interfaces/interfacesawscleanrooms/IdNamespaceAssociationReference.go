package interfacesawscleanrooms


// A reference to a IdNamespaceAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idNamespaceAssociationReference := &IdNamespaceAssociationReference{
//   	IdNamespaceAssociationArn: jsii.String("idNamespaceAssociationArn"),
//   	IdNamespaceAssociationIdentifier: jsii.String("idNamespaceAssociationIdentifier"),
//   	MembershipIdentifier: jsii.String("membershipIdentifier"),
//   }
//
type IdNamespaceAssociationReference struct {
	// The ARN of the IdNamespaceAssociation resource.
	IdNamespaceAssociationArn *string `field:"required" json:"idNamespaceAssociationArn" yaml:"idNamespaceAssociationArn"`
	// The IdNamespaceAssociationIdentifier of the IdNamespaceAssociation resource.
	IdNamespaceAssociationIdentifier *string `field:"required" json:"idNamespaceAssociationIdentifier" yaml:"idNamespaceAssociationIdentifier"`
	// The MembershipIdentifier of the IdNamespaceAssociation resource.
	MembershipIdentifier *string `field:"required" json:"membershipIdentifier" yaml:"membershipIdentifier"`
}

