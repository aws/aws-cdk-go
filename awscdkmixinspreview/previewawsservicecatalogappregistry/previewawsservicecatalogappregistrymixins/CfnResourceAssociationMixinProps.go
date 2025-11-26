package previewawsservicecatalogappregistrymixins


// Properties for CfnResourceAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnResourceAssociationMixinProps := &CfnResourceAssociationMixinProps{
//   	Application: jsii.String("application"),
//   	Resource: jsii.String("resource"),
//   	ResourceType: jsii.String("resourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-resourceassociation.html
//
type CfnResourceAssociationMixinProps struct {
	// The name or ID of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-resourceassociation.html#cfn-servicecatalogappregistry-resourceassociation-application
	//
	Application *string `field:"optional" json:"application" yaml:"application"`
	// The name or ID of the resource of which the application will be associated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-resourceassociation.html#cfn-servicecatalogappregistry-resourceassociation-resource
	//
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
	// The type of resource of which the application will be associated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-resourceassociation.html#cfn-servicecatalogappregistry-resourceassociation-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

