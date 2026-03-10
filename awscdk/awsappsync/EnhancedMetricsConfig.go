package awsappsync


// Enhanced metrics configuration for AppSync.
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
type EnhancedMetricsConfig struct {
	// Controls how data source metrics will be emitted to CloudWatch.
	DataSourceLevelMetricsBehavior DataSourceLevelMetricsBehavior `field:"required" json:"dataSourceLevelMetricsBehavior" yaml:"dataSourceLevelMetricsBehavior"`
	// Controls how operation metrics will be emitted to CloudWatch.
	OperationLevelMetricsConfig OperationLevelMetricsConfig `field:"required" json:"operationLevelMetricsConfig" yaml:"operationLevelMetricsConfig"`
	// Controls how resolver metrics will be emitted to CloudWatch.
	ResolverLevelMetricsBehavior ResolverLevelMetricsBehavior `field:"required" json:"resolverLevelMetricsBehavior" yaml:"resolverLevelMetricsBehavior"`
}

