package awsservicecatalogappregistry


// Properties for defining a `CfnResourceAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnResourceAssociationProps := &cfnResourceAssociationProps{
//   	application: jsii.String("application"),
//   	resource: jsii.String("resource"),
//   	resourceType: jsii.String("resourceType"),
//   }
//
type CfnResourceAssociationProps struct {
	// The name or ID of the application.
	Application *string `field:"required" json:"application" yaml:"application"`
	// The name or ID of the resource of which the application will be associated.
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The type of resource of which the application will be associated.
	//
	// Possible values: CFN_STACK.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

