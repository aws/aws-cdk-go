package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnServerlessCacheSnapshot`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServerlessCacheSnapshotProps := &CfnServerlessCacheSnapshotProps{
//   	ServerlessCacheName: jsii.String("serverlessCacheName"),
//   	ServerlessCacheSnapshotName: jsii.String("serverlessCacheSnapshotName"),
//
//   	// the properties below are optional
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscachesnapshot.html
//
type CfnServerlessCacheSnapshotProps struct {
	// The name of an existing serverless cache.
	//
	// The snapshot is created from this cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscachesnapshot.html#cfn-elasticache-serverlesscachesnapshot-serverlesscachename
	//
	ServerlessCacheName *string `field:"required" json:"serverlessCacheName" yaml:"serverlessCacheName"`
	// The name of the serverless cache snapshot.
	//
	// Must be unique for the customer account. This value is stored as a lowercase string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscachesnapshot.html#cfn-elasticache-serverlesscachesnapshot-serverlesscachesnapshotname
	//
	ServerlessCacheSnapshotName *string `field:"required" json:"serverlessCacheSnapshotName" yaml:"serverlessCacheSnapshotName"`
	// The Amazon Resource Name (ARN) of the AWS KMS key used to encrypt the snapshot.
	//
	// Provide the key ARN: the resource returns the key ARN on read, so supplying a bare key ID or alias for this createOnly property may be reported as drift by CloudFormation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscachesnapshot.html#cfn-elasticache-serverlesscachesnapshot-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A list of tags to be added to the serverless cache snapshot resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscachesnapshot.html#cfn-elasticache-serverlesscachesnapshot-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

