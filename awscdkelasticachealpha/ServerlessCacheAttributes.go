package awscdkelasticachealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Attributes that can be specified when importing a ServerlessCache.
//
// Example:
//   var securityGroup SecurityGroup
//
//
//   importedServerlessCache := elasticache.ServerlessCache_FromServerlessCacheAttributes(this, jsii.String("ImportedServerlessCache"), &ServerlessCacheAttributes{
//   	ServerlessCacheName: jsii.String("my-serverless-cache"),
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   })
//
// Experimental.
type ServerlessCacheAttributes struct {
	// The ARNs of backups restored in the cache.
	// Default: - backups are unknown.
	//
	// Experimental.
	BackupArnsToRestore *[]*string `field:"optional" json:"backupArnsToRestore" yaml:"backupArnsToRestore"`
	// The cache engine used by this cache.
	// Default: - engine type is unknown.
	//
	// Experimental.
	Engine CacheEngine `field:"optional" json:"engine" yaml:"engine"`
	// The KMS key used for encryption.
	// Default: - encryption key is unknown.
	//
	// Experimental.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The security groups associated with this cache.
	// Default: - security groups are unknown.
	//
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The ARN of the serverless cache.
	//
	// One of `serverlessCacheName` or `serverlessCacheArn` is required.
	// Default: - derived from serverlessCacheName.
	//
	// Experimental.
	ServerlessCacheArn *string `field:"optional" json:"serverlessCacheArn" yaml:"serverlessCacheArn"`
	// The name of the serverless cache.
	//
	// One of `serverlessCacheName` or `serverlessCacheArn` is required.
	// Default: - derived from serverlessCacheArn.
	//
	// Experimental.
	ServerlessCacheName *string `field:"optional" json:"serverlessCacheName" yaml:"serverlessCacheName"`
	// The subnets this cache is deployed in.
	// Default: - subnets are unknown.
	//
	// Experimental.
	Subnets *[]awsec2.ISubnet `field:"optional" json:"subnets" yaml:"subnets"`
	// The user group associated with this cache.
	// Default: - user group is unknown.
	//
	// Experimental.
	UserGroup IUserGroup `field:"optional" json:"userGroup" yaml:"userGroup"`
	// The VPC this cache is deployed in.
	// Default: - VPC is unknown.
	//
	// Experimental.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

