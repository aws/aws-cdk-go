package interfacesawsiottwinmaker


// A reference to a SyncJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syncJobReference := &SyncJobReference{
//   	SyncJobArn: jsii.String("syncJobArn"),
//   	SyncSource: jsii.String("syncSource"),
//   	WorkspaceId: jsii.String("workspaceId"),
//   }
//
type SyncJobReference struct {
	// The ARN of the SyncJob resource.
	SyncJobArn *string `field:"required" json:"syncJobArn" yaml:"syncJobArn"`
	// The SyncSource of the SyncJob resource.
	SyncSource *string `field:"required" json:"syncSource" yaml:"syncSource"`
	// The WorkspaceId of the SyncJob resource.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

