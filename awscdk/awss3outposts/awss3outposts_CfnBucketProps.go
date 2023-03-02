package awss3outposts

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnBucket`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var filter interface{}
//
//   cfnBucketProps := &cfnBucketProps{
//   	bucketName: jsii.String("bucketName"),
//   	outpostId: jsii.String("outpostId"),
//
//   	// the properties below are optional
//   	lifecycleConfiguration: &lifecycleConfigurationProperty{
//   		rules: []interface{}{
//   			&ruleProperty{
//   				status: jsii.String("status"),
//
//   				// the properties below are optional
//   				abortIncompleteMultipartUpload: &abortIncompleteMultipartUploadProperty{
//   					daysAfterInitiation: jsii.Number(123),
//   				},
//   				expirationDate: jsii.String("expirationDate"),
//   				expirationInDays: jsii.Number(123),
//   				filter: filter,
//   				id: jsii.String("id"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBucketProps struct {
	// A name for the S3 on Outposts bucket.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. The bucket name must contain only lowercase letters, numbers, periods (.), and dashes (-) and must follow [Amazon S3 bucket restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/BucketRestrictions.html) . For more information, see [Bucket naming rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/BucketRestrictions.html#bucketnamingrules) .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The ID of the Outpost of the specified bucket.
	OutpostId *string `field:"required" json:"outpostId" yaml:"outpostId"`
	// Creates a new lifecycle configuration for the S3 on Outposts bucket or replaces an existing lifecycle configuration.
	//
	// Outposts buckets only support lifecycle configurations that delete/expire objects after a certain period of time and abort incomplete multipart uploads.
	LifecycleConfiguration interface{} `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// Sets the tags for an S3 on Outposts bucket. For more information, see [Using Amazon S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) .
	//
	// Use tags to organize your AWS bill to reflect your own cost structure. To do this, sign up to get your AWS account bill with tag key values included. Then, to see the cost of combined resources, organize your billing information according to resources with the same tag key values. For example, you can tag several resources with a specific application name, and then organize your billing information to see the total cost of that application across several services. For more information, see [Cost allocation and tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) .
	//
	// > Within a bucket, if you add a tag that has the same key as an existing tag, the new value overwrites the old value. For more information, see [Using cost allocation and bucket tags](https://docs.aws.amazon.com/AmazonS3/latest/userguide/CostAllocTagging.html) .
	//
	// To use this resource, you must have permissions to perform the `s3-outposts:PutBucketTagging` . The S3 on Outposts bucket owner has this permission by default and can grant this permission to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations](https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources) and [Managing access permissions to your Amazon S3 resources](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

