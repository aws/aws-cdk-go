package previewawsappsyncmixins


// Describes an enhanced metrics configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   enhancedMetricsConfigProperty := &EnhancedMetricsConfigProperty{
//   	DataSourceLevelMetricsBehavior: jsii.String("dataSourceLevelMetricsBehavior"),
//   	OperationLevelMetricsConfig: jsii.String("operationLevelMetricsConfig"),
//   	ResolverLevelMetricsBehavior: jsii.String("resolverLevelMetricsBehavior"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-enhancedmetricsconfig.html
//
type CfnGraphQLApiPropsMixin_EnhancedMetricsConfigProperty struct {
	// Controls how data source metrics will be emitted to CloudWatch. Data source metrics include:.
	//
	// - *Requests* : The number of invocations that occured during a request.
	// - *Latency* : The time to complete a data source invocation.
	// - *Errors* : The number of errors that occurred during a data source invocation.
	//
	// These metrics can be emitted to CloudWatch per data source or for all data sources in the request. Metrics will be recorded by API ID and data source name. `dataSourceLevelMetricsBehavior` accepts one of these values at a time:
	//
	// - `FULL_REQUEST_DATA_SOURCE_METRICS` : Records and emits metric data for all data sources in the request.
	// - `PER_DATA_SOURCE_METRICS` : Records and emits metric data for data sources that have the `MetricsConfig` value set to `ENABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-enhancedmetricsconfig.html#cfn-appsync-graphqlapi-enhancedmetricsconfig-datasourcelevelmetricsbehavior
	//
	DataSourceLevelMetricsBehavior *string `field:"optional" json:"dataSourceLevelMetricsBehavior" yaml:"dataSourceLevelMetricsBehavior"`
	// Controls how operation metrics will be emitted to CloudWatch. Operation metrics include:.
	//
	// - *Requests* : The number of times a specified GraphQL operation was called.
	// - *GraphQL errors* : The number of GraphQL errors that occurred during a specified GraphQL operation.
	//
	// Metrics will be recorded by API ID and operation name. You can set the value to `ENABLED` or `DISABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-enhancedmetricsconfig.html#cfn-appsync-graphqlapi-enhancedmetricsconfig-operationlevelmetricsconfig
	//
	OperationLevelMetricsConfig *string `field:"optional" json:"operationLevelMetricsConfig" yaml:"operationLevelMetricsConfig"`
	// Controls how resolver metrics will be emitted to CloudWatch. Resolver metrics include:.
	//
	// - *GraphQL errors* : The number of GraphQL errors that occurred.
	// - *Requests* : The number of invocations that occurred during a request.
	// - *Latency* : The time to complete a resolver invocation.
	// - *Cache hits* : The number of cache hits during a request.
	// - *Cache misses* : The number of cache misses during a request.
	//
	// These metrics can be emitted to CloudWatch per resolver or for all resolvers in the request. Metrics will be recorded by API ID and resolver name. `resolverLevelMetricsBehavior` accepts one of these values at a time:
	//
	// - `FULL_REQUEST_RESOLVER_METRICS` : Records and emits metric data for all resolvers in the request.
	// - `PER_RESOLVER_METRICS` : Records and emits metric data for resolvers that have the `MetricsConfig` value set to `ENABLED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-graphqlapi-enhancedmetricsconfig.html#cfn-appsync-graphqlapi-enhancedmetricsconfig-resolverlevelmetricsbehavior
	//
	ResolverLevelMetricsBehavior *string `field:"optional" json:"resolverLevelMetricsBehavior" yaml:"resolverLevelMetricsBehavior"`
}

