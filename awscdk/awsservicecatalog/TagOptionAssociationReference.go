package awsservicecatalog


// A reference to a TagOptionAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagOptionAssociationReference := &TagOptionAssociationReference{
//   	TagOptionAssociationId: jsii.String("tagOptionAssociationId"),
//   }
//
type TagOptionAssociationReference struct {
	// The Id of the TagOptionAssociation resource.
	TagOptionAssociationId *string `field:"required" json:"tagOptionAssociationId" yaml:"tagOptionAssociationId"`
}

