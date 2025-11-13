package interfacesawsiottwinmaker


// A reference to a ComponentType resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   componentTypeReference := &ComponentTypeReference{
//   	ComponentTypeArn: jsii.String("componentTypeArn"),
//   	ComponentTypeId: jsii.String("componentTypeId"),
//   	WorkspaceId: jsii.String("workspaceId"),
//   }
//
type ComponentTypeReference struct {
	// The ARN of the ComponentType resource.
	ComponentTypeArn *string `field:"required" json:"componentTypeArn" yaml:"componentTypeArn"`
	// The ComponentTypeId of the ComponentType resource.
	ComponentTypeId *string `field:"required" json:"componentTypeId" yaml:"componentTypeId"`
	// The WorkspaceId of the ComponentType resource.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

