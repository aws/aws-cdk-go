package awscdkelasticachealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for defining a ServerlessCache.
//
// Example:
//   var vpc Vpc
//
//
//   serverlessCache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
//   	Engine: elasticache.CacheEngine_VALKEY_LATEST,
//   	Backup: &BackupSettings{
//   		// set a backup name before deleting a cache
//   		BackupNameBeforeDeletion: jsii.String("my-final-backup-name"),
//   	},
//   	Vpc: Vpc,
//   })
//
// Experimental.
type ServerlessCacheProps struct {
	// The VPC to place the cache in.
	// Experimental.
	Vpc awsec2.IVpc `field:"required" json:"vpc" yaml:"vpc"`
	// Backup configuration.
	// Default: - No backups configured.
	//
	// Experimental.
	Backup *BackupSettings `field:"optional" json:"backup" yaml:"backup"`
	// Usage limits for the cache.
	// Default: - No usage limits.
	//
	// Experimental.
	CacheUsageLimits *CacheUsageLimitsProperty `field:"optional" json:"cacheUsageLimits" yaml:"cacheUsageLimits"`
	// A description for the cache.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The cache engine combined with the version Enum options: VALKEY_DEFAULT, VALKEY_7, VALKEY_8, REDIS_DEFAULT, MEMCACHED_DEFAULT The default options bring the latest versions available.
	// Default: when not provided, the default engine would be Valkey, latest version available (VALKEY_DEFAULT).
	//
	// Experimental.
	Engine CacheEngine `field:"optional" json:"engine" yaml:"engine"`
	// KMS key for encryption.
	// Default: - Service managed encryption (AWS owned KMS key).
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Security groups for the cache.
	// Default: - A new security group is created.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Name for the serverless cache.
	// Default: automatically generated name by Resource.
	//
	// Experimental.
	ServerlessCacheName *string `field:"optional" json:"serverlessCacheName" yaml:"serverlessCacheName"`
	// User group for access control.
	// Default: - No user group.
	//
	// Experimental.
	UserGroup IUserGroup `field:"optional" json:"userGroup" yaml:"userGroup"`
	// Which subnets to place the cache in.
	// Default: - Private subnets with egress.
	//
	// Experimental.
	VpcSubnets *awsec2.SubnetSelection `field:"optional" json:"vpcSubnets" yaml:"vpcSubnets"`
}

