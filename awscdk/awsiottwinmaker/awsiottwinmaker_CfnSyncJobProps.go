package awsiottwinmaker


// Properties for defining a `CfnSyncJob`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSyncJobProps := &cfnSyncJobProps{
//   	syncRole: jsii.String("syncRole"),
//   	syncSource: jsii.String("syncSource"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSyncJobProps struct {
	// `AWS::IoTTwinMaker::SyncJob.SyncRole`.
	SyncRole *string `field:"required" json:"syncRole" yaml:"syncRole"`
	// `AWS::IoTTwinMaker::SyncJob.SyncSource`.
	SyncSource *string `field:"required" json:"syncSource" yaml:"syncSource"`
	// `AWS::IoTTwinMaker::SyncJob.WorkspaceId`.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// `AWS::IoTTwinMaker::SyncJob.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

