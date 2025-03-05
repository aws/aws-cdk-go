package awsservicecatalogappregistry


// Properties for defining a `CfnResourceAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceAssociationProps := &CfnResourceAssociationProps{
//   	Application: jsii.String("application"),
//   	Resource: jsii.String("resource"),
//   	ResourceType: jsii.String("resourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-resourceassociation.html
//
type CfnResourceAssociationProps struct {
	// The name or ID of the application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-resourceassociation.html#cfn-servicecatalogappregistry-resourceassociation-application
	//
	Application *string `field:"required" json:"application" yaml:"application"`
	// The name or ID of the resource of which the application will be associated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-resourceassociation.html#cfn-servicecatalogappregistry-resourceassociation-resource
	//
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The type of resource of which the application will be associated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalogappregistry-resourceassociation.html#cfn-servicecatalogappregistry-resourceassociation-resourcetype
	//
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

