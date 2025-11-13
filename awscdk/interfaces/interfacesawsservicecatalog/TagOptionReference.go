package interfacesawsservicecatalog


// A reference to a TagOption resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagOptionReference := &TagOptionReference{
//   	TagOptionId: jsii.String("tagOptionId"),
//   }
//
type TagOptionReference struct {
	// The Id of the TagOption resource.
	TagOptionId *string `field:"required" json:"tagOptionId" yaml:"tagOptionId"`
}

