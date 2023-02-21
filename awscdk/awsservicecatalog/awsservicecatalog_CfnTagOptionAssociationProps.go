package awsservicecatalog


// Properties for defining a `CfnTagOptionAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnTagOptionAssociationProps := &cfnTagOptionAssociationProps{
//   	resourceId: jsii.String("resourceId"),
//   	tagOptionId: jsii.String("tagOptionId"),
//   }
//
type CfnTagOptionAssociationProps struct {
	// The resource identifier.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The TagOption identifier.
	TagOptionId *string `field:"required" json:"tagOptionId" yaml:"tagOptionId"`
}

