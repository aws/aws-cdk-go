package awsgamelift


// The location in Amazon S3 where build or script files can be stored for access by Amazon GameLift.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3LocationProperty := &s3LocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	objectVersion: jsii.String("objectVersion"),
//   }
//
type CfnScript_S3LocationProperty struct {
	// An Amazon S3 bucket identifier. This is the name of the S3 bucket.
	//
	// > GameLift currently does not support uploading from Amazon S3 buckets with names that contain a dot (.).
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the zip file that contains the build files or script files.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The Amazon Resource Name ( [ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html) ) for an IAM role that allows Amazon GameLift to access the S3 bucket.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The version of the file, if object versioning is turned on for the bucket.
	//
	// Amazon GameLift uses this information when retrieving files from an S3 bucket that you own. Use this parameter to specify a specific version of the file. If not set, the latest version of the file is retrieved.
	ObjectVersion *string `field:"optional" json:"objectVersion" yaml:"objectVersion"`
}

