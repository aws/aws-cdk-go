package awsomics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRunCache`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRunCacheProps := &CfnRunCacheProps{
//   	CacheBehavior: jsii.String("cacheBehavior"),
//   	CacheBucketOwnerId: jsii.String("cacheBucketOwnerId"),
//   	CacheS3Location: jsii.String("cacheS3Location"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-runcache.html
//
type CfnRunCacheProps struct {
	// The default cache behavior for runs using this cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-runcache.html#cfn-omics-runcache-cachebehavior
	//
	CacheBehavior *string `field:"optional" json:"cacheBehavior" yaml:"cacheBehavior"`
	// The AWS account ID of the expected owner of the S3 bucket for the run cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-runcache.html#cfn-omics-runcache-cachebucketownerid
	//
	CacheBucketOwnerId *string `field:"optional" json:"cacheBucketOwnerId" yaml:"cacheBucketOwnerId"`
	// The S3 location for storing the cached task outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-runcache.html#cfn-omics-runcache-caches3location
	//
	CacheS3Location *string `field:"optional" json:"cacheS3Location" yaml:"cacheS3Location"`
	// A description of the run cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-runcache.html#cfn-omics-runcache-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the run cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-runcache.html#cfn-omics-runcache-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags for the run cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-runcache.html#cfn-omics-runcache-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

