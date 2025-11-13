package interfacesawsservicecatalogappregistry


// A reference to a ResourceAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceAssociationReference := &ResourceAssociationReference{
//   	ApplicationArn: jsii.String("applicationArn"),
//   	ResourceArn: jsii.String("resourceArn"),
//   	ResourceType: jsii.String("resourceType"),
//   }
//
type ResourceAssociationReference struct {
	// The ApplicationArn of the ResourceAssociation resource.
	ApplicationArn *string `field:"required" json:"applicationArn" yaml:"applicationArn"`
	// The ResourceArn of the ResourceAssociation resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The ResourceType of the ResourceAssociation resource.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

