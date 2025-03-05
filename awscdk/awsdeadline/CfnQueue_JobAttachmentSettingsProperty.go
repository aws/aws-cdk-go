package awsdeadline


// The job attachment settings.
//
// These are the Amazon S3 bucket name and the Amazon S3 prefix.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobAttachmentSettingsProperty := &JobAttachmentSettingsProperty{
//   	RootPrefix: jsii.String("rootPrefix"),
//   	S3BucketName: jsii.String("s3BucketName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-jobattachmentsettings.html
//
type CfnQueue_JobAttachmentSettingsProperty struct {
	// The root prefix.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-jobattachmentsettings.html#cfn-deadline-queue-jobattachmentsettings-rootprefix
	//
	RootPrefix *string `field:"required" json:"rootPrefix" yaml:"rootPrefix"`
	// The Amazon S3 bucket name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-queue-jobattachmentsettings.html#cfn-deadline-queue-jobattachmentsettings-s3bucketname
	//
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
}

