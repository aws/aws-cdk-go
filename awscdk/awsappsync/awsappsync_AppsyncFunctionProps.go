package awsappsync


// the CDK properties for AppSync Functions.
//
// Example:
//   var api graphqlApi
//
//
//   appsyncFunction := appsync.NewAppsyncFunction(this, jsii.String("function"), &AppsyncFunctionProps{
//   	Name: jsii.String("appsync_function"),
//   	Api: Api,
//   	DataSource: api.AddNoneDataSource(jsii.String("none")),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
//   })
//
// Experimental.
type AppsyncFunctionProps struct {
	// the name of the AppSync Function.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// the description for this AppSync Function.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// the request mapping template for the AppSync Function.
	// Experimental.
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// the response mapping template for the AppSync Function.
	// Experimental.
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// the GraphQL Api linked to this AppSync Function.
	// Experimental.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// the data source linked to this AppSync Function.
	// Experimental.
	DataSource BaseDataSource `field:"required" json:"dataSource" yaml:"dataSource"`
}

