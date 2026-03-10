package awsappsync


// Controls how operation metrics will be emitted to CloudWatch.
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
type OperationLevelMetricsConfig string

const (
	// Sends operation metrics to CloudWatch.
	OperationLevelMetricsConfig_ENABLED OperationLevelMetricsConfig = "ENABLED"
	// Does not send operation metrics to CloudWatch.
	OperationLevelMetricsConfig_DISABLED OperationLevelMetricsConfig = "DISABLED"
)

