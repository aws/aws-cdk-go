package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocationS3`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationS3Props := &cfnLocationS3Props{
//   	s3BucketArn: jsii.String("s3BucketArn"),
//   	s3Config: &s3ConfigProperty{
//   		bucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   	},
//
//   	// the properties below are optional
//   	s3StorageClass: jsii.String("s3StorageClass"),
//   	subdirectory: jsii.String("subdirectory"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationS3Props struct {
	// The ARN of the Amazon S3 bucket.
	S3BucketArn *string `field:"required" json:"s3BucketArn" yaml:"s3BucketArn"`
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that is used to access an Amazon S3 bucket.
	//
	// For detailed information about using such a role, see [Creating a Location for Amazon S3](https://docs.aws.amazon.com/datasync/latest/userguide/working-with-locations.html#create-s3-location) in the *AWS DataSync User Guide* .
	S3Config interface{} `field:"required" json:"s3Config" yaml:"s3Config"`
	// The Amazon S3 storage class that you want to store your files in when this location is used as a task destination.
	//
	// For buckets in AWS Regions , the storage class defaults to S3 Standard.
	//
	// For more information about S3 storage classes, see [Amazon S3 Storage Classes](https://docs.aws.amazon.com/s3/storage-classes/) . Some storage classes have behaviors that can affect your S3 storage costs. For detailed information, see [Considerations When Working with Amazon S3 Storage Classes in DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes) .
	S3StorageClass *string `field:"optional" json:"s3StorageClass" yaml:"s3StorageClass"`
	// A subdirectory in the Amazon S3 bucket.
	//
	// This subdirectory in Amazon S3 is used to read data from the S3 source location or write data to the S3 destination.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// The key-value pair that represents the tag that you want to add to the location.
	//
	// The value can be an empty string. We recommend using tags to name your resources.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

