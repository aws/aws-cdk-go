package awselasticache

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnServerlessCache`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServerlessCacheProps := &CfnServerlessCacheProps{
//   	Engine: jsii.String("engine"),
//   	ServerlessCacheName: jsii.String("serverlessCacheName"),
//
//   	// the properties below are optional
//   	CacheUsageLimits: &CacheUsageLimitsProperty{
//   		DataStorage: &DataStorageProperty{
//   			Maximum: jsii.Number(123),
//   			Unit: jsii.String("unit"),
//   		},
//   		EcpuPerSecond: &ECPUPerSecondProperty{
//   			Maximum: jsii.Number(123),
//   		},
//   	},
//   	DailySnapshotTime: jsii.String("dailySnapshotTime"),
//   	Description: jsii.String("description"),
//   	FinalSnapshotName: jsii.String("finalSnapshotName"),
//   	KmsKeyId: jsii.String("kmsKeyId"),
//   	MajorEngineVersion: jsii.String("majorEngineVersion"),
//   	SecurityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	SnapshotArnsToRestore: []*string{
//   		jsii.String("snapshotArnsToRestore"),
//   	},
//   	SnapshotRetentionLimit: jsii.Number(123),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	UserGroupId: jsii.String("userGroupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html
//
type CfnServerlessCacheProps struct {
	// The engine the serverless cache is compatible with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-engine
	//
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The unique identifier of the serverless cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-serverlesscachename
	//
	ServerlessCacheName *string `field:"required" json:"serverlessCacheName" yaml:"serverlessCacheName"`
	// The cache usage limit for the serverless cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-cacheusagelimits
	//
	CacheUsageLimits interface{} `field:"optional" json:"cacheUsageLimits" yaml:"cacheUsageLimits"`
	// The daily time that a cache snapshot will be created.
	//
	// Default is NULL, i.e. snapshots will not be created at a specific time on a daily basis. Available for Redis only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-dailysnapshottime
	//
	DailySnapshotTime *string `field:"optional" json:"dailySnapshotTime" yaml:"dailySnapshotTime"`
	// A description of the serverless cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the final snapshot taken of a cache before the cache is deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-finalsnapshotname
	//
	FinalSnapshotName *string `field:"optional" json:"finalSnapshotName" yaml:"finalSnapshotName"`
	// The ID of the AWS Key Management Service (KMS) key that is used to encrypt data at rest in the serverless cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-kmskeyid
	//
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The version number of the engine the serverless cache is compatible with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-majorengineversion
	//
	MajorEngineVersion *string `field:"optional" json:"majorEngineVersion" yaml:"majorEngineVersion"`
	// The IDs of the EC2 security groups associated with the serverless cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-securitygroupids
	//
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ARN of the snapshot from which to restore data into the new cache.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-snapshotarnstorestore
	//
	SnapshotArnsToRestore *[]*string `field:"optional" json:"snapshotArnsToRestore" yaml:"snapshotArnsToRestore"`
	// The current setting for the number of serverless cache snapshots the system will retain.
	//
	// Available for Redis only.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-snapshotretentionlimit
	//
	SnapshotRetentionLimit *float64 `field:"optional" json:"snapshotRetentionLimit" yaml:"snapshotRetentionLimit"`
	// If no subnet IDs are given and your VPC is in SFO, then ElastiCache will select 2 default subnets across AZs in your VPC.
	//
	// For all other Regions, if no subnet IDs are given then ElastiCache will select 3 default subnets across AZs in your default VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// A list of tags to be added to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The identifier of the user group associated with the serverless cache.
	//
	// Available for Redis only. Default is NULL.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-serverlesscache.html#cfn-elasticache-serverlesscache-usergroupid
	//
	UserGroupId *string `field:"optional" json:"userGroupId" yaml:"userGroupId"`
}

