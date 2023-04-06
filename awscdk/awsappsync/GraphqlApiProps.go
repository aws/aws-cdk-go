package awsappsync


// Properties for an AppSync GraphQL API.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &GraphqlApiProps{
//   	Name: jsii.String("api"),
//   	Schema: appsync.SchemaFile_FromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
//   })
//
//   httpDs := api.AddHttpDataSource(jsii.String("ds"), jsii.String("https://states.amazonaws.com"), &HttpDataSourceOptions{
//   	Name: jsii.String("httpDsWithStepF"),
//   	Description: jsii.String("from appsync to StepFunctions Workflow"),
//   	AuthorizationConfig: &AwsIamConfig{
//   		SigningRegion: jsii.String("us-east-1"),
//   		SigningServiceName: jsii.String("states"),
//   	},
//   })
//
//   httpDs.CreateResolver(jsii.String("MutationCallStepFunctionResolver"), &BaseResolverProps{
//   	TypeName: jsii.String("Mutation"),
//   	FieldName: jsii.String("callStepFunction"),
//   	RequestMappingTemplate: appsync.MappingTemplate_FromFile(jsii.String("request.vtl")),
//   	ResponseMappingTemplate: appsync.MappingTemplate_*FromFile(jsii.String("response.vtl")),
//   })
//
type GraphqlApiProps struct {
	// the name of the GraphQL API.
	Name *string `field:"required" json:"name" yaml:"name"`
	// GraphQL schema definition. Specify how you want to define your schema.
	//
	// Schema.fromFile(filePath: string) allows schema definition through schema.graphql file
	Schema ISchema `field:"required" json:"schema" yaml:"schema"`
	// Optional authorization configuration.
	AuthorizationConfig *AuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The domain name configuration for the GraphQL API.
	//
	// The Route 53 hosted zone and CName DNS record must be configured in addition to this setting to
	// enable custom domain URL.
	DomainName *DomainOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Logging configuration for this api.
	LogConfig *LogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// A flag indicating whether or not X-Ray tracing is enabled for the GraphQL API.
	XrayEnabled *bool `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

