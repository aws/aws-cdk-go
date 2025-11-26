package previewawsdeadlinemixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnQueuePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnQueueMixinProps := &CfnQueueMixinProps{
//   	AllowedStorageProfileIds: []*string{
//   		jsii.String("allowedStorageProfileIds"),
//   	},
//   	DefaultBudgetAction: jsii.String("defaultBudgetAction"),
//   	Description: jsii.String("description"),
//   	DisplayName: jsii.String("displayName"),
//   	FarmId: jsii.String("farmId"),
//   	JobAttachmentSettings: &JobAttachmentSettingsProperty{
//   		RootPrefix: jsii.String("rootPrefix"),
//   		S3BucketName: jsii.String("s3BucketName"),
//   	},
//   	JobRunAsUser: &JobRunAsUserProperty{
//   		Posix: &PosixUserProperty{
//   			Group: jsii.String("group"),
//   			User: jsii.String("user"),
//   		},
//   		RunAs: jsii.String("runAs"),
//   		Windows: &WindowsUserProperty{
//   			PasswordArn: jsii.String("passwordArn"),
//   			User: jsii.String("user"),
//   		},
//   	},
//   	RequiredFileSystemLocationNames: []*string{
//   		jsii.String("requiredFileSystemLocationNames"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html
//
type CfnQueueMixinProps struct {
	// The identifiers of the storage profiles that this queue can use to share assets between workers using different operating systems.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html#cfn-deadline-queue-allowedstorageprofileids
	//
	AllowedStorageProfileIds *[]*string `field:"optional" json:"allowedStorageProfileIds" yaml:"allowedStorageProfileIds"`
	// The default action taken on a queue summary if a budget wasn't configured.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html#cfn-deadline-queue-defaultbudgetaction
	//
	// Default: - "NONE".
	//
	DefaultBudgetAction *string `field:"optional" json:"defaultBudgetAction" yaml:"defaultBudgetAction"`
	// A description of the queue that helps identify what the queue is used for.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html#cfn-deadline-queue-description
	//
	// Default: - "".
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the queue summary to update.
	//
	// > This field can store any content. Escape or encode this content before displaying it on a webpage or any other system that might interpret the content of this field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html#cfn-deadline-queue-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The farm ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html#cfn-deadline-queue-farmid
	//
	FarmId *string `field:"optional" json:"farmId" yaml:"farmId"`
	// The job attachment settings.
	//
	// These are the Amazon S3 bucket name and the Amazon S3 prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html#cfn-deadline-queue-jobattachmentsettings
	//
	JobAttachmentSettings interface{} `field:"optional" json:"jobAttachmentSettings" yaml:"jobAttachmentSettings"`
	// Identifies the user for a job.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html#cfn-deadline-queue-jobrunasuser
	//
	JobRunAsUser interface{} `field:"optional" json:"jobRunAsUser" yaml:"jobRunAsUser"`
	// The file system location that the queue uses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html#cfn-deadline-queue-requiredfilesystemlocationnames
	//
	RequiredFileSystemLocationNames *[]*string `field:"optional" json:"requiredFileSystemLocationNames" yaml:"requiredFileSystemLocationNames"`
	// The Amazon Resource Name (ARN) of the IAM role that workers use when running jobs in this queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html#cfn-deadline-queue-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The tags to add to your queue.
	//
	// Each tag consists of a tag key and a tag value. Tag keys and values are both required, but tag values can be empty strings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-deadline-queue.html#cfn-deadline-queue-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

