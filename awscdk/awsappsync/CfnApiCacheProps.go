package awsappsync


// Properties for defining a `CfnApiCache`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiCacheProps := &CfnApiCacheProps{
//   	ApiCachingBehavior: jsii.String("apiCachingBehavior"),
//   	ApiId: jsii.String("apiId"),
//   	Ttl: jsii.Number(123),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AtRestEncryptionEnabled: jsii.Boolean(false),
//   	HealthMetricsConfig: jsii.String("healthMetricsConfig"),
//   	TransitEncryptionEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html
//
type CfnApiCacheProps struct {
	// Caching behavior.
	//
	// - *FULL_REQUEST_CACHING* : All requests from the same user are cached. Individual resolvers are automatically cached. All API calls will try to return responses from the cache.
	// - *PER_RESOLVER_CACHING* : Individual resolvers that you specify are cached.
	// - *OPERATION_LEVEL_CACHING* : Full requests are cached together and returned without executing resolvers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-apicachingbehavior
	//
	ApiCachingBehavior *string `field:"required" json:"apiCachingBehavior" yaml:"apiCachingBehavior"`
	// The GraphQL API ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-apiid
	//
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// TTL in seconds for cache entries.
	//
	// Valid values are 1–3,600 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-ttl
	//
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
	// The cache instance type. Valid values are.
	//
	// - `SMALL`
	// - `MEDIUM`
	// - `LARGE`
	// - `XLARGE`
	// - `LARGE_2X`
	// - `LARGE_4X`
	// - `LARGE_8X` (not available in all regions)
	// - `LARGE_12X`
	//
	// Historically, instance types were identified by an EC2-style value. As of July 2020, this is deprecated, and the generic identifiers above should be used.
	//
	// The following legacy instance types are available, but their use is discouraged:
	//
	// - *T2_SMALL* : A t2.small instance type.
	// - *T2_MEDIUM* : A t2.medium instance type.
	// - *R4_LARGE* : A r4.large instance type.
	// - *R4_XLARGE* : A r4.xlarge instance type.
	// - *R4_2XLARGE* : A r4.2xlarge instance type.
	// - *R4_4XLARGE* : A r4.4xlarge instance type.
	// - *R4_8XLARGE* : A r4.8xlarge instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// *This parameter has been deprecated* .
	//
	// At-rest encryption flag for cache. You cannot update this setting after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-atrestencryptionenabled
	//
	AtRestEncryptionEnabled interface{} `field:"optional" json:"atRestEncryptionEnabled" yaml:"atRestEncryptionEnabled"`
	// Controls how cache health metrics will be emitted to CloudWatch. Cache health metrics include:.
	//
	// - *NetworkBandwidthOutAllowanceExceeded* : The network packets dropped because the throughput exceeded the aggregated bandwidth limit. This is useful for diagnosing bottlenecks in a cache configuration.
	// - *EngineCPUUtilization* : The CPU utilization (percentage) allocated to the Redis process. This is useful for diagnosing bottlenecks in a cache configuration.
	//
	// Metrics will be recorded by API ID. You can set the value to `ENABLED` or `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-healthmetricsconfig
	//
	HealthMetricsConfig *string `field:"optional" json:"healthMetricsConfig" yaml:"healthMetricsConfig"`
	// *This parameter has been deprecated* .
	//
	// Transit encryption flag when connecting to cache. You cannot update this setting after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-transitencryptionenabled
	//
	TransitEncryptionEnabled interface{} `field:"optional" json:"transitEncryptionEnabled" yaml:"transitEncryptionEnabled"`
}

