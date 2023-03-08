package awsappsync


// Optional configuration for Http data sources.
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
type HttpDataSourceOptions struct {
	// The description of the data source.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source, overrides the id given by cdk.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The authorization config in case the HTTP endpoint requires authorization.
	AuthorizationConfig *AwsIamConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

