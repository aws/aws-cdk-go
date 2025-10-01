package awsgamelift


// A reference to a ContainerGroupDefinition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerGroupDefinitionReference := &ContainerGroupDefinitionReference{
//   	ContainerGroupDefinitionArn: jsii.String("containerGroupDefinitionArn"),
//   	ContainerGroupDefinitionName: jsii.String("containerGroupDefinitionName"),
//   }
//
type ContainerGroupDefinitionReference struct {
	// The ARN of the ContainerGroupDefinition resource.
	ContainerGroupDefinitionArn *string `field:"required" json:"containerGroupDefinitionArn" yaml:"containerGroupDefinitionArn"`
	// The Name of the ContainerGroupDefinition resource.
	ContainerGroupDefinitionName *string `field:"required" json:"containerGroupDefinitionName" yaml:"containerGroupDefinitionName"`
}

