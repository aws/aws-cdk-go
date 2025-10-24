package awss3express

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDirectoryBucket`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDirectoryBucketProps := &CfnDirectoryBucketProps{
//   	DataRedundancy: jsii.String("dataRedundancy"),
//   	LocationName: jsii.String("locationName"),
//
//   	// the properties below are optional
//   	BucketEncryption: &BucketEncryptionProperty{
//   		ServerSideEncryptionConfiguration: []interface{}{
//   			&ServerSideEncryptionRuleProperty{
//   				BucketKeyEnabled: jsii.Boolean(false),
//   				ServerSideEncryptionByDefault: &ServerSideEncryptionByDefaultProperty{
//   					SseAlgorithm: jsii.String("sseAlgorithm"),
//
//   					// the properties below are optional
//   					KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   				},
//   			},
//   		},
//   	},
//   	BucketName: jsii.String("bucketName"),
//   	LifecycleConfiguration: &LifecycleConfigurationProperty{
//   		Rules: []interface{}{
//   			&RuleProperty{
//   				Status: jsii.String("status"),
//
//   				// the properties below are optional
//   				AbortIncompleteMultipartUpload: &AbortIncompleteMultipartUploadProperty{
//   					DaysAfterInitiation: jsii.Number(123),
//   				},
//   				ExpirationInDays: jsii.Number(123),
//   				Id: jsii.String("id"),
//   				ObjectSizeGreaterThan: jsii.String("objectSizeGreaterThan"),
//   				ObjectSizeLessThan: jsii.String("objectSizeLessThan"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html
//
type CfnDirectoryBucketProps struct {
	// The number of Zone (Availability Zone or Local Zone) that's used for redundancy for the bucket.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html#cfn-s3express-directorybucket-dataredundancy
	//
	DataRedundancy *string `field:"required" json:"dataRedundancy" yaml:"dataRedundancy"`
	// The name of the location where the bucket will be created.
	//
	// For directory buckets, the name of the location is the Zone ID of the Availability Zone (AZ) or Local Zone (LZ) where the bucket will be created. An example AZ ID value is `usw2-az1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html#cfn-s3express-directorybucket-locationname
	//
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
	// Specifies default encryption for a bucket using server-side encryption with Amazon S3 managed keys (SSE-S3) or AWS KMS keys (SSE-KMS).
	//
	// For information about default encryption for directory buckets, see [Setting and monitoring default encryption for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-bucket-encryption.html) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html#cfn-s3express-directorybucket-bucketencryption
	//
	BucketEncryption interface{} `field:"optional" json:"bucketEncryption" yaml:"bucketEncryption"`
	// A name for the bucket.
	//
	// The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Zone (Availability Zone or Local Zone). The bucket name must also follow the format `*bucket_base_name* -- *zone_id* --x-s3` (for example, `*bucket_base_name* -- *usw2-az1* --x-s3` ). If you don't specify a name, AWS CloudFormation generates a unique ID and uses that ID for the bucket name. For information about bucket naming restrictions, see [Directory bucket naming rules](https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-bucket-naming-rules.html) in the *Amazon S3 User Guide* .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you need to replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html#cfn-s3express-directorybucket-bucketname
	//
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Container for lifecycle rules. You can add as many as 1000 rules.
	//
	// For more information see, [Creating and managing a lifecycle configuration for directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-objects-lifecycle.html          ) in the *Amazon S3 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html#cfn-s3express-directorybucket-lifecycleconfiguration
	//
	LifecycleConfiguration interface{} `field:"optional" json:"lifecycleConfiguration" yaml:"lifecycleConfiguration"`
	// An array of tags that you can apply to the S3 directory bucket.
	//
	// Tags are key-value pairs of metadata used to categorize and organize your buckets, track costs, and control access. For more information, see [Using tags with directory buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-tagging.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3express-directorybucket.html#cfn-s3express-directorybucket-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

