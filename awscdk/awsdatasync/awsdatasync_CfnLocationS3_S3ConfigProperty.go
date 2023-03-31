package awsdatasync


// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role used to access an Amazon S3 bucket.
//
// For detailed information about using such a role, see [Creating a Location for Amazon S3](https://docs.aws.amazon.com/datasync/latest/userguide/working-with-locations.html#create-s3-location) in the *AWS DataSync User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigProperty := &s3ConfigProperty{
//   	bucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   }
//
type CfnLocationS3_S3ConfigProperty struct {
	// The ARN of the IAM role for accessing the S3 bucket.
	BucketAccessRoleArn *string `field:"required" json:"bucketAccessRoleArn" yaml:"bucketAccessRoleArn"`
}

