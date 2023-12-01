package awselasticache


// The usage limits for storage and ElastiCache Processing Units for the cache.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cacheUsageLimitsProperty := &CacheUsageLimitsProperty{
//   	DataStorage: &DataStorageProperty{
//   		Maximum: jsii.Number(123),
//   		Unit: jsii.String("unit"),
//   	},
//   	EcpuPerSecond: &ECPUPerSecondProperty{
//   		Maximum: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-cacheusagelimits.html
//
type CfnServerlessCache_CacheUsageLimitsProperty struct {
	// The maximum data storage limit in the cache, expressed in Gigabytes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-cacheusagelimits.html#cfn-elasticache-serverlesscache-cacheusagelimits-datastorage
	//
	DataStorage interface{} `field:"optional" json:"dataStorage" yaml:"dataStorage"`
	// The number of ElastiCache Processing Units (ECPU) the cache can consume per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-cacheusagelimits.html#cfn-elasticache-serverlesscache-cacheusagelimits-ecpupersecond
	//
	EcpuPerSecond interface{} `field:"optional" json:"ecpuPerSecond" yaml:"ecpuPerSecond"`
}

