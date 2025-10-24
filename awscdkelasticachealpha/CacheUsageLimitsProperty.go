package awscdkelasticachealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Usage limits configuration for ServerlessCache.
//
// Example:
//   var vpc Vpc
//
//
//   serverlessCache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
//   	Engine: elasticache.CacheEngine_VALKEY_LATEST,
//   	Vpc: Vpc,
//   	CacheUsageLimits: &CacheUsageLimitsProperty{
//   		// cache data storage limits (GB)
//   		DataStorageMinimumSize: awscdk.Size_Gibibytes(jsii.Number(2)),
//   		 // minimum: 1GB
//   		DataStorageMaximumSize: awscdk.Size_*Gibibytes(jsii.Number(3)),
//   		 // maximum: 5000GB
//   		// rate limits (ECPU/second)
//   		RequestRateLimitMinimum: jsii.Number(1000),
//   		 // minimum: 1000
//   		RequestRateLimitMaximum: jsii.Number(10000),
//   	},
//   })
//
// Experimental.
type CacheUsageLimitsProperty struct {
	// Maximum data storage size (5000 GB).
	// Default: - No maximum limit.
	//
	// Experimental.
	DataStorageMaximumSize awscdk.Size `field:"optional" json:"dataStorageMaximumSize" yaml:"dataStorageMaximumSize"`
	// Minimum data storage size (1 GB).
	// Default: - No minimum limit.
	//
	// Experimental.
	DataStorageMinimumSize awscdk.Size `field:"optional" json:"dataStorageMinimumSize" yaml:"dataStorageMinimumSize"`
	// Maximum request rate limit (15000000 ECPUs per second).
	// Default: - No maximum limit.
	//
	// Experimental.
	RequestRateLimitMaximum *float64 `field:"optional" json:"requestRateLimitMaximum" yaml:"requestRateLimitMaximum"`
	// Minimum request rate limit (1000 ECPUs per second).
	// Default: - No minimum limit.
	//
	// Experimental.
	RequestRateLimitMinimum *float64 `field:"optional" json:"requestRateLimitMinimum" yaml:"requestRateLimitMinimum"`
}

