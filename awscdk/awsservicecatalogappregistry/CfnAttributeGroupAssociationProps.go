package awsservicecatalogappregistry


// Properties for defining a `CfnAttributeGroupAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAttributeGroupAssociationProps := &CfnAttributeGroupAssociationProps{
//   	Application: jsii.String("application"),
//   	AttributeGroup: jsii.String("attributeGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroupassociation.html
//
type CfnAttributeGroupAssociationProps struct {
	// The name or ID of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroupassociation.html#cfn-servicecatalogappregistry-attributegroupassociation-application
	//
	Application *string `field:"required" json:"application" yaml:"application"`
	// The name or ID of the attribute group that holds the attributes to describe the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroupassociation.html#cfn-servicecatalogappregistry-attributegroupassociation-attributegroup
	//
	AttributeGroup *string `field:"required" json:"attributeGroup" yaml:"attributeGroup"`
}

