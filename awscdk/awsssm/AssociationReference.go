package awsssm


// A reference to a Association resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   associationReference := &AssociationReference{
//   	AssociationId: jsii.String("associationId"),
//   }
//
type AssociationReference struct {
	// The AssociationId of the Association resource.
	AssociationId *string `field:"required" json:"associationId" yaml:"associationId"`
}

