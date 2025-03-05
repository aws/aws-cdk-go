package awsdatasync


// Specifies the manifest that you want AWS DataSync to use and where it's hosted.
//
// For more information and configuration examples, see [Specifying what DataSync transfers by using a manifest](https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &SourceProperty{
//   	S3: &ManifestConfigSourceS3Property{
//   		BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   		ManifestObjectPath: jsii.String("manifestObjectPath"),
//   		ManifestObjectVersionId: jsii.String("manifestObjectVersionId"),
//   		S3BucketArn: jsii.String("s3BucketArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-source.html
//
type CfnTask_SourceProperty struct {
	// Specifies the S3 bucket where you're hosting your manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-source.html#cfn-datasync-task-source-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}

