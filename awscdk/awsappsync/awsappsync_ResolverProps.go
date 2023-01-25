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
// Experimental.
type ResolverProps struct {
	// name of the GraphQL field in the given type this resolver is attached to.
	// Experimental.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
	// name of the GraphQL type this resolver is attached to.
	// Experimental.
	TypeName *string `field:"required" json:"typeName" yaml:"typeName"`
	// The caching configuration for this resolver.
	// Experimental.
	CachingConfig *CachingConfig `field:"optional" json:"cachingConfig" yaml:"cachingConfig"`
	// configuration of the pipeline resolver.
	// Experimental.
	PipelineConfig *[]IAppsyncFunction `field:"optional" json:"pipelineConfig" yaml:"pipelineConfig"`
	// The request mapping template for this resolver.
	// Experimental.
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// The response mapping template for this resolver.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The data source this resolver is using.
	// Experimental.
	DataSource BaseDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
	// The API this resolver is attached to.
	// Experimental.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
}

