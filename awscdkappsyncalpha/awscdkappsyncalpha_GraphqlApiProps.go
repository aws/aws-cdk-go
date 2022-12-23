// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// Properties for an AppSync GraphQL API.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
//   	name: jsii.String("api"),
//   	schema: appsync.schemaFile.fromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
//   })
//
//   httpDs := api.addHttpDataSource(jsii.String("ds"), jsii.String("https://states.amazonaws.com"), &httpDataSourceOptions{
//   	name: jsii.String("httpDsWithStepF"),
//   	description: jsii.String("from appsync to StepFunctions Workflow"),
//   	authorizationConfig: &awsIamConfig{
//   		signingRegion: jsii.String("us-east-1"),
//   		signingServiceName: jsii.String("states"),
//   	},
//   })
//
//   httpDs.createResolver(jsii.String("MutationCallStepFunctionResolver"), &baseResolverProps{
//   	typeName: jsii.String("Mutation"),
//   	fieldName: jsii.String("callStepFunction"),
//   	requestMappingTemplate: appsync.mappingTemplate.fromFile(jsii.String("request.vtl")),
//   	responseMappingTemplate: appsync.*mappingTemplate.fromFile(jsii.String("response.vtl")),
//   })
//
// Experimental.
type GraphqlApiProps struct {
	// the name of the GraphQL API.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// GraphQL schema definition. Specify how you want to define your schema.
	//
	// Schema.fromFile(filePath: string) allows schema definition through schema.graphql file
	// Experimental.
	Schema ISchema `field:"required" json:"schema" yaml:"schema"`
	// Optional authorization configuration.
	// Experimental.
	AuthorizationConfig *AuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// The domain name configuration for the GraphQL API.
	//
	// The Route 53 hosted zone and CName DNS record must be configured in addition to this setting to
	// enable custom domain URL.
	// Experimental.
	DomainName *DomainOptions `field:"optional" json:"domainName" yaml:"domainName"`
	// Logging configuration for this api.
	// Experimental.
	LogConfig *LogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// A flag indicating whether or not X-Ray tracing is enabled for the GraphQL API.
	// Experimental.
	XrayEnabled *bool `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

