package awsservicecatalogappregistry


// Properties for defining a `CfnAttributeGroupAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAttributeGroupAssociationProps := &cfnAttributeGroupAssociationProps{
//   	application: jsii.String("application"),
//   	attributeGroup: jsii.String("attributeGroup"),
//   }
//
type CfnAttributeGroupAssociationProps struct {
	// The name or ID of the application.
	Application *string `field:"required" json:"application" yaml:"application"`
	// The name or ID of the attribute group that holds the attributes to describe the application.
	AttributeGroup *string `field:"required" json:"attributeGroup" yaml:"attributeGroup"`
}

