package interfacesawscases


// A reference to a RelatedItem resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relatedItemReference := &RelatedItemReference{
//   	RelatedItemArn: jsii.String("relatedItemArn"),
//   }
//
type RelatedItemReference struct {
	// The RelatedItemArn of the RelatedItem resource.
	RelatedItemArn *string `field:"required" json:"relatedItemArn" yaml:"relatedItemArn"`
}

