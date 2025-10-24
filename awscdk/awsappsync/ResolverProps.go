package awsappsync


// Additional property for an AppSync resolver for GraphQL API reference.
//
// Example:
//   var api GraphqlApi
//   var appsyncFunction AppsyncFunction
//
//
//   pipelineResolver := appsync.NewResolver(this, jsii.String("pipeline"), &ResolverProps{
//   	Api: Api,
//   	DataSource: api.AddNoneDataSource(jsii.String("none")),
//   	TypeName: jsii.String("typeName"),
//   	FieldName: jsii.String("fieldName"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("beforeRequest.vtl")),
//   	PipelineConfig: []IAppsyncFunction{
//   		appsyncFunction,
//   	},
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("afterResponse.vtl")),
//   })
//
type ResolverProps struct {
	// name of the GraphQL field in the given type this resolver is attached to.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// name of the GraphQL type this resolver is attached to.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The caching configuration for this resolver.
	// Default: - No caching configuration.
	//
	CachingConfig *CachingConfig `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// The function code.
	// Default: - no code is used.
	//
	Code Code `field:"optional" json:"code" yaml:"code"`
	// The maximum number of elements per batch, when using batch invoke.
	// Default: - No max batch size.
	//
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// configuration of the pipeline resolver.
	// Default: - no pipeline resolver configuration
	// An empty array | undefined sets resolver to be of kind, unit.
	//
	PipelineConfig *[]IAppsyncFunction `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template for this resolver.
	// Default: - No mapping template.
	//
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	// Default: - No mapping template.
	//
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The functions runtime.
	// Default: - no function runtime, VTL mapping templates used.
	//
	Runtime FunctionRuntime `field:"optional" json:"runtime" yaml:"runtime"`
	// The data source this resolver is using.
	// Default: - No datasource.
	//
	DataSource BaseDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The API this resolver is attached to.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
}

