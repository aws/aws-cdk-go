package previewawsiottwinmakermixins


// Properties for CfnSyncJobPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSyncJobMixinProps := &CfnSyncJobMixinProps{
//   	SyncRole: jsii.String("syncRole"),
//   	SyncSource: jsii.String("syncSource"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	WorkspaceId: jsii.String("workspaceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-syncjob.html
//
type CfnSyncJobMixinProps struct {
	// The SyncJob IAM role.
	//
	// This IAM role is used by the sync job to read from the syncSource, and create, update or delete the corresponding resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-syncjob.html#cfn-iottwinmaker-syncjob-syncrole
	//
	SyncRole *string `field:"optional" json:"syncRole" yaml:"syncRole"`
	// The sync source.
	//
	// > Currently the only supported syncSoucre is `SITEWISE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-syncjob.html#cfn-iottwinmaker-syncjob-syncsource
	//
	SyncSource *string `field:"optional" json:"syncSource" yaml:"syncSource"`
	// Metadata you can use to manage the SyncJob.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-syncjob.html#cfn-iottwinmaker-syncjob-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the workspace that contains the sync job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-syncjob.html#cfn-iottwinmaker-syncjob-workspaceid
	//
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

