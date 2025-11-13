package interfacesawsiottwinmaker


// A reference to a Entity resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityReference := &EntityReference{
//   	EntityArn: jsii.String("entityArn"),
//   	EntityId: jsii.String("entityId"),
//   	WorkspaceId: jsii.String("workspaceId"),
//   }
//
type EntityReference struct {
	// The ARN of the Entity resource.
	EntityArn *string `field:"required" json:"entityArn" yaml:"entityArn"`
	// The EntityId of the Entity resource.
	EntityId *string `field:"required" json:"entityId" yaml:"entityId"`
	// The WorkspaceId of the Entity resource.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

