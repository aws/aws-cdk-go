package awsgreengrass


// A reference to a ResourceDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceDefinitionReference := &ResourceDefinitionReference{
//   	ResourceDefinitionArn: jsii.String("resourceDefinitionArn"),
//   	ResourceDefinitionId: jsii.String("resourceDefinitionId"),
//   }
//
type ResourceDefinitionReference struct {
	// The ARN of the ResourceDefinition resource.
	ResourceDefinitionArn *string `field:"required" json:"resourceDefinitionArn" yaml:"resourceDefinitionArn"`
	// The Id of the ResourceDefinition resource.
	ResourceDefinitionId *string `field:"required" json:"resourceDefinitionId" yaml:"resourceDefinitionId"`
}

