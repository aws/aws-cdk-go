package awsappstream


// A reference to a StackUserAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stackUserAssociationReference := &StackUserAssociationReference{
//   	StackUserAssociationId: jsii.String("stackUserAssociationId"),
//   }
//
type StackUserAssociationReference struct {
	// The Id of the StackUserAssociation resource.
	StackUserAssociationId *string `field:"required" json:"stackUserAssociationId" yaml:"stackUserAssociationId"`
}

