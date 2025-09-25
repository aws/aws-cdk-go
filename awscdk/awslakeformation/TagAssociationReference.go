package awslakeformation


// A reference to a TagAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagAssociationReference := &TagAssociationReference{
//   	ResourceIdentifier: jsii.String("resourceIdentifier"),
//   	TagsIdentifier: jsii.String("tagsIdentifier"),
//   }
//
type TagAssociationReference struct {
	// The ResourceIdentifier of the TagAssociation resource.
	ResourceIdentifier *string `field:"required" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// The TagsIdentifier of the TagAssociation resource.
	TagsIdentifier *string `field:"required" json:"tagsIdentifier" yaml:"tagsIdentifier"`
}

