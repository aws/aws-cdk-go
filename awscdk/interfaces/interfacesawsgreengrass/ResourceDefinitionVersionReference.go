package interfacesawsgreengrass


// A reference to a ResourceDefinitionVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDefinitionVersionReference := &ResourceDefinitionVersionReference{
//   	ResourceDefinitionVersionId: jsii.String("resourceDefinitionVersionId"),
//   }
//
type ResourceDefinitionVersionReference struct {
	// The Id of the ResourceDefinitionVersion resource.
	ResourceDefinitionVersionId *string `field:"required" json:"resourceDefinitionVersionId" yaml:"resourceDefinitionVersionId"`
}

