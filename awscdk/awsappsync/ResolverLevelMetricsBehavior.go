package awsappsync


// Controls how resolver metrics will be emitted to CloudWatch.
//
// Example:
//   schema := appsync.NewSchemaFile(&SchemaProps{
//   	FilePath: jsii.String("mySchemaFile"),
//   })
//   appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("myApi"),
//   	Definition: appsync.Definition_FromSchema(schema),
//   	EnhancedMetricsConfig: &EnhancedMetricsConfig{
//   		DataSourceLevelMetricsBehavior: appsync.DataSourceLevelMetricsBehavior_FULL_REQUEST_DATA_SOURCE_METRICS,
//   		OperationLevelMetricsConfig: appsync.OperationLevelMetricsConfig_ENABLED,
//   		ResolverLevelMetricsBehavior: appsync.ResolverLevelMetricsBehavior_FULL_REQUEST_RESOLVER_METRICS,
//   	},
//   })
//
type ResolverLevelMetricsBehavior string

const (
	// Records and emits metric data for all resolvers in the request.
	ResolverLevelMetricsBehavior_FULL_REQUEST_RESOLVER_METRICS ResolverLevelMetricsBehavior = "FULL_REQUEST_RESOLVER_METRICS"
	// Records and emits metric data for resolvers that have the MetricsConfig value set to ENABLED.
	ResolverLevelMetricsBehavior_PER_RESOLVER_METRICS ResolverLevelMetricsBehavior = "PER_RESOLVER_METRICS"
)

