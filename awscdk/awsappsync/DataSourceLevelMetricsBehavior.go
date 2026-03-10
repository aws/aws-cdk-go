package awsappsync


// Controls how data source metrics will be emitted to CloudWatch.
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
type DataSourceLevelMetricsBehavior string

const (
	// Records and emits metric data for all data sources in the request.
	DataSourceLevelMetricsBehavior_FULL_REQUEST_DATA_SOURCE_METRICS DataSourceLevelMetricsBehavior = "FULL_REQUEST_DATA_SOURCE_METRICS"
	// Records and emits metric data for data sources that have the MetricsConfig value set to ENABLED.
	DataSourceLevelMetricsBehavior_PER_DATA_SOURCE_METRICS DataSourceLevelMetricsBehavior = "PER_DATA_SOURCE_METRICS"
)

