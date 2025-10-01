package awssecurityhub


// A reference to a PolicyAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   policyAssociationReference := &PolicyAssociationReference{
//   	AssociationIdentifier: jsii.String("associationIdentifier"),
//   }
//
type PolicyAssociationReference struct {
	// The AssociationIdentifier of the PolicyAssociation resource.
	AssociationIdentifier *string `field:"required" json:"associationIdentifier" yaml:"associationIdentifier"`
}

