package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnLocationS3`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationS3Props := &CfnLocationS3Props{
//   	S3Config: &S3ConfigProperty{
//   		BucketAccessRoleArn: jsii.String("bucketAccessRoleArn"),
//   	},
//
//   	// the properties below are optional
//   	S3BucketArn: jsii.String("s3BucketArn"),
//   	S3StorageClass: jsii.String("s3StorageClass"),
//   	Subdirectory: jsii.String("subdirectory"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locations3.html
//
type CfnLocationS3Props struct {
	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that is used to access an Amazon S3 bucket.
	//
	// For detailed information about using such a role, see [Creating a Location for Amazon S3](https://docs.aws.amazon.com/datasync/latest/userguide/working-with-locations.html#create-s3-location) in the *AWS DataSync User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locations3.html#cfn-datasync-locations3-s3config
	//
	S3Config interface{} `field:"required" json:"s3Config" yaml:"s3Config"`
	// The ARN of the Amazon S3 bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locations3.html#cfn-datasync-locations3-s3bucketarn
	//
	S3BucketArn interface{} `field:"optional" json:"s3BucketArn" yaml:"s3BucketArn"`
	// The Amazon S3 storage class that you want to store your files in when this location is used as a task destination.
	//
	// For buckets in AWS Regions , the storage class defaults to S3 Standard.
	//
	// For more information about S3 storage classes, see [Amazon S3 Storage Classes](https://docs.aws.amazon.com/s3/storage-classes/) . Some storage classes have behaviors that can affect your S3 storage costs. For detailed information, see [Considerations When Working with Amazon S3 Storage Classes in DataSync](https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locations3.html#cfn-datasync-locations3-s3storageclass
	//
	// Default: - "STANDARD".
	//
	S3StorageClass *string `field:"optional" json:"s3StorageClass" yaml:"s3StorageClass"`
	// Specifies a prefix in the S3 bucket that DataSync reads from or writes to (depending on whether the bucket is a source or destination location).
	//
	// > DataSync can't transfer objects with a prefix that begins with a slash ( `/` ) or includes `//` , `/./` , or `/../` patterns. For example:
	// >
	// > - `/photos`
	// > - `photos//2006/January`
	// > - `photos/./2006/February`
	// > - `photos/../2006/March`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locations3.html#cfn-datasync-locations3-subdirectory
	//
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies labels that help you categorize, filter, and search for your AWS resources.
	//
	// We recommend creating at least a name tag for your transfer location.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-locations3.html#cfn-datasync-locations3-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

