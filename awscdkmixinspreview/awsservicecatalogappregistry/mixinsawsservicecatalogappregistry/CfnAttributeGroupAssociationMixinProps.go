package mixinsawsservicecatalogappregistry


// Properties for CfnAttributeGroupAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAttributeGroupAssociationMixinProps := &CfnAttributeGroupAssociationMixinProps{
//   	Application: jsii.String("application"),
//   	AttributeGroup: jsii.String("attributeGroup"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroupassociation.html
//
type CfnAttributeGroupAssociationMixinProps struct {
	// The name or ID of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroupassociation.html#cfn-servicecatalogappregistry-attributegroupassociation-application
	//
	Application *string `field:"optional" json:"application" yaml:"application"`
	// The name or ID of the attribute group which holds the attributes that describe the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-attributegroupassociation.html#cfn-servicecatalogappregistry-attributegroupassociation-attributegroup
	//
	AttributeGroup *string `field:"optional" json:"attributeGroup" yaml:"attributeGroup"`
}

