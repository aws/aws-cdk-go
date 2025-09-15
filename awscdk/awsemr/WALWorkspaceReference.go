package awsemr


// A reference to a WALWorkspace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   wALWorkspaceReference := &WALWorkspaceReference{
//   	WalWorkspaceName: jsii.String("walWorkspaceName"),
//   }
//
type WALWorkspaceReference struct {
	// The WALWorkspaceName of the WALWorkspace resource.
	WalWorkspaceName *string `field:"required" json:"walWorkspaceName" yaml:"walWorkspaceName"`
}

