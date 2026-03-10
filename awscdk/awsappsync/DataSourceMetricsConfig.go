package awsappsync


// Enum for enhanced data source metrics for specified data sources.
//
// Example:
//   schema := appsync.NewSchemaFile(&SchemaProps{
//   	FilePath: jsii.String("mySchemaFile"),
//   })
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("myApi"),
//   	Definition: appsync.Definition_FromSchema(schema),
//   	EnhancedMetricsConfig: &EnhancedMetricsConfig{
//   		DataSourceLevelMetricsBehavior: appsync.DataSourceLevelMetricsBehavior_PER_DATA_SOURCE_METRICS,
//   		OperationLevelMetricsConfig: appsync.OperationLevelMetricsConfig_ENABLED,
//   		ResolverLevelMetricsBehavior: appsync.ResolverLevelMetricsBehavior_PER_RESOLVER_METRICS,
//   	},
//   })
//
//   noneDS := api.AddNoneDataSource(jsii.String("none"), &DataSourceOptions{
//   	MetricsConfig: appsync.DataSourceMetricsConfig_ENABLED,
//   })
//
//   noneDS.CreateResolver(jsii.String("noneResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("addDemoMetricsConfig"),
//   	MetricsConfig: appsync.ResolverMetricsConfig_ENABLED,
//   })
//
type DataSourceMetricsConfig string

const (
	// Enables enhanced data source metrics for specified data sources.
	DataSourceMetricsConfig_ENABLED DataSourceMetricsConfig = "ENABLED"
	// Disables enhanced data source metrics for specified data sources.
	DataSourceMetricsConfig_DISABLED DataSourceMetricsConfig = "DISABLED"
)

