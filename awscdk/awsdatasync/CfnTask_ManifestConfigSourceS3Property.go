package awsdatasync


// Specifies the S3 bucket where you're hosting the manifest that you want AWS DataSync to use.
//
// For more information and configuration examples, see [Specifying what DataSync transfers by using a manifest](https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   manifestConfigSourceS3Property := &ManifestConfigSourceS3Property{
//   	BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   	ManifestObjectPath: jsii.String("manifestObjectPath"),
//   	ManifestObjectVersionId: jsii.String("manifestObjectVersionId"),
//   	S3BucketArn: jsii.String("s3BucketArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-manifestconfigsources3.html
//
type CfnTask_ManifestConfigSourceS3Property struct {
	// Specifies the AWS Identity and Access Management (IAM) role that allows DataSync to access your manifest.
	//
	// For more information, see [Providing DataSync access to your manifest](https://docs.aws.amazon.com/datasync/latest/userguide/transferring-with-manifest.html#transferring-with-manifest-access) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-manifestconfigsources3.html#cfn-datasync-task-manifestconfigsources3-bucketaccessrolearn
	//
	BucketAccessRoleArn *string `field:"optional" json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
	// Specifies the Amazon S3 object key of your manifest.
	//
	// This can include a prefix (for example, `prefix/my-manifest.csv` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-manifestconfigsources3.html#cfn-datasync-task-manifestconfigsources3-manifestobjectpath
	//
	ManifestObjectPath *string `field:"optional" json:"manifestObjectPath" yaml:"manifestObjectPath"`
	// Specifies the object version ID of the manifest that you want DataSync to use.
	//
	// If you don't set this, DataSync uses the latest version of the object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-manifestconfigsources3.html#cfn-datasync-task-manifestconfigsources3-manifestobjectversionid
	//
	ManifestObjectVersionId *string `field:"optional" json:"manifestObjectVersionId" yaml:"manifestObjectVersionId"`
	// Specifies the Amazon Resource Name (ARN) of the S3 bucket where you're hosting your manifest.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-task-manifestconfigsources3.html#cfn-datasync-task-manifestconfigsources3-s3bucketarn
	//
	S3BucketArn *string `field:"optional" json:"s3BucketArn" yaml:"s3BucketArn"`
}

