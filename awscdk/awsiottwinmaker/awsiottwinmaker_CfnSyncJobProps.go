package awsiottwinmaker


// Properties for defining a `CfnSyncJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSyncJobProps := &CfnSyncJobProps{
//   	SyncRole: jsii.String("syncRole"),
//   	SyncSource: jsii.String("syncSource"),
//   	WorkspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSyncJobProps struct {
	// The SyncJob IAM role.
	//
	// This IAM role is used by the sync job to read from the syncSource, and create, update or delete the corresponding resources.
	SyncRole *string `field:"required" json:"syncRole" yaml:"syncRole"`
	// The sync source.
	//
	// > Currently the only supported syncSoucre is `SITEWISE` .
	SyncSource *string `field:"required" json:"syncSource" yaml:"syncSource"`
	// The ID of the workspace that contains the sync job.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Metadata you can use to manage the SyncJob.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

