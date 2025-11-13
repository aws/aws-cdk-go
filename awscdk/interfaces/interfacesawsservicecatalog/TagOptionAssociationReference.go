package interfacesawsservicecatalog


// A reference to a TagOptionAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagOptionAssociationReference := &TagOptionAssociationReference{
//   	ResourceId: jsii.String("resourceId"),
//   	TagOptionId: jsii.String("tagOptionId"),
//   }
//
type TagOptionAssociationReference struct {
	// The ResourceId of the TagOptionAssociation resource.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The TagOptionId of the TagOptionAssociation resource.
	TagOptionId *string `field:"required" json:"tagOptionId" yaml:"tagOptionId"`
}

