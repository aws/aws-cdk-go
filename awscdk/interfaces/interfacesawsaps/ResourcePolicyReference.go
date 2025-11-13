package interfacesawsaps


// A reference to a ResourcePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourcePolicyReference := &ResourcePolicyReference{
//   	WorkspaceArn: jsii.String("workspaceArn"),
//   }
//
type ResourcePolicyReference struct {
	// The WorkspaceArn of the ResourcePolicy resource.
	WorkspaceArn *string `field:"required" json:"workspaceArn" yaml:"workspaceArn"`
}

