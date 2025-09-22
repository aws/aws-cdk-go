package awsservicecatalogappregistry


// A reference to a AttributeGroupAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeGroupAssociationReference := &AttributeGroupAssociationReference{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	AttributeGroupArn: jsii.String("attributeGroupArn"),
//   }
//
type AttributeGroupAssociationReference struct {
	// The ApplicationArn of the AttributeGroupAssociation resource.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The AttributeGroupArn of the AttributeGroupAssociation resource.
	AttributeGroupArn *string `field:"required" json:"attributeGroupArn" yaml:"attributeGroupArn"`
}

