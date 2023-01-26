package awsappsync


// Additional property for an AppSync resolver for GraphQL API reference.
//
// Example:
//   var api graphqlApi
//   var appsyncFunction appsyncFunction
//
//
//   pipelineResolver := appsync.NewResolver(this, jsii.String("pipeline"), &resolverProps{
//   	api: api,
//   	dataSource: api.addNoneDataSource(jsii.String("none")),
//   	typeName: jsii.String("typeName"),
//   	fieldName: jsii.String("fieldName"),
//   	requestMappingTemplate: appsync.mappingTemplate.fromFile(jsii.String("beforeRequest.vtl")),
//   	pipelineConfig: []iAppsyncFunction{
//   		appsyncFunction,
//   	},
//   	responseMappingTemplate: appsync.*mappingTemplate.fromFile(jsii.String("afterResponse.vtl")),
//   })
//
type ResolverProps struct {
	// name of the GraphQL field in the given type this resolver is attached to.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// name of the GraphQL type this resolver is attached to.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The caching configuration for this resolver.
	CachingConfig *CachingConfig `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// The function code.
	Code Code `field:"optional" json:"code" yaml:"code"`
	// The maximum number of elements per batch, when using batch invoke.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// configuration of the pipeline resolver.
	PipelineConfig *[]IAppsyncFunction `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template for this resolver.
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The functions runtime.
	Runtime FunctionRuntime `field:"optional" json:"runtime" yaml:"runtime"`
	// The data source this resolver is using.
	DataSource BaseDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The API this resolver is attached to.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
}

