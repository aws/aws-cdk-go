package awsappsync


// the CDK properties for AppSync Functions.
//
// Example:
//   var api GraphqlApi
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
type AppsyncFunctionProps struct {
	// the name of the AppSync Function.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The function code.
	// Default: - no code is used.
	//
	Code Code `field:"optional" json:"code" yaml:"code"`
	// the description for this AppSync Function.
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The maximum number of resolver request inputs that will be sent to a single AWS Lambda function in a BatchInvoke operation.
	//
	// Can only be set when using LambdaDataSource.
	// Default: - No max batch size.
	//
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// the request mapping template for the AppSync Function.
	// Default: - no request mapping template.
	//
	RequestMappingTemplate MappingTemplate `field:"optional" json:"requestMappingTemplate" yaml:"requestMappingTemplate"`
	// the response mapping template for the AppSync Function.
	// Default: - no response mapping template.
	//
	ResponseMappingTemplate MappingTemplate `field:"optional" json:"responseMappingTemplate" yaml:"responseMappingTemplate"`
	// The functions runtime.
	// Default: - no function runtime, VTL mapping templates used.
	//
	Runtime FunctionRuntime `field:"optional" json:"runtime" yaml:"runtime"`
	// the GraphQL Api linked to this AppSync Function.
	Api IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// the data source linked to this AppSync Function.
	DataSource BaseDataSource `field:"required" json:"dataSource" yaml:"dataSource"`
}

