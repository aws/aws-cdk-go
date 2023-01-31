package awsappsync


// Optional configuration for Http data sources.
//
// Example:
//   api := appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
//   	name: jsii.String("api"),
//   	schema: appsync.schema.fromAsset(path.join(__dirname, jsii.String("schema.graphql"))),
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
//   httpDs.createResolver(&baseResolverProps{
//   	typeName: jsii.String("Mutation"),
//   	fieldName: jsii.String("callStepFunction"),
//   	requestMappingTemplate: appsync.mappingTemplate.fromFile(jsii.String("request.vtl")),
//   	responseMappingTemplate: appsync.*mappingTemplate.fromFile(jsii.String("response.vtl")),
//   })
//
// Experimental.
type HttpDataSourceOptions struct {
	// The description of the data source.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the data source, overrides the id given by cdk.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The authorization config in case the HTTP endpoint requires authorization.
	// Experimental.
	AuthorizationConfig *AwsIamConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

