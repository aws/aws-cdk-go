package awsdatasync


// Specifies the Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that DataSync uses to access your S3 bucket.
//
// For more information, see [Providing DataSync access to S3 buckets](https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#create-s3-location-access) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigProperty := &S3ConfigProperty{
//   	BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locations3-s3config.html
//
type CfnLocationS3_S3ConfigProperty struct {
	// Specifies the ARN of the IAM role that DataSync uses to access your S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datasync-locations3-s3config.html#cfn-datasync-locations3-s3config-bucketaccessrolearn
	//
	BucketAccessRoleArn *string `field:"required" json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
}

