package interfacesawsworkspacesinstances


// A reference to a WorkspaceInstance resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workspaceInstanceReference := &WorkspaceInstanceReference{
//   	WorkspaceInstanceId: jsii.String("workspaceInstanceId"),
//   }
//
type WorkspaceInstanceReference struct {
	// The WorkspaceInstanceId of the WorkspaceInstance resource.
	WorkspaceInstanceId *string `field:"required" json:"workspaceInstanceId" yaml:"workspaceInstanceId"`
}

