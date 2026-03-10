package awsappsync


// The options for configuring a schema from an existing file.
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
type SchemaProps struct {
	// The file path for the schema.
	//
	// When this option is
	// configured, then the schema will be generated from an
	// existing file from disk.
	FilePath *string `field:"required" json:"filePath" yaml:"filePath"`
}

