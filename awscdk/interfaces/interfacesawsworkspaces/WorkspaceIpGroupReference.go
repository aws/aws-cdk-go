package interfacesawsworkspaces


// A reference to a WorkspaceIpGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workspaceIpGroupReference := &WorkspaceIpGroupReference{
//   	WorkspaceIpGroupArn: jsii.String("workspaceIpGroupArn"),
//   }
//
type WorkspaceIpGroupReference struct {
	// The Arn of the WorkspaceIpGroup resource.
	WorkspaceIpGroupArn *string `field:"required" json:"workspaceIpGroupArn" yaml:"workspaceIpGroupArn"`
}

