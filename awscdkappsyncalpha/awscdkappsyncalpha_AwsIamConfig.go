// The CDK Construct Library for AWS::AppSync
package awscdkappsyncalpha


// The authorization config in case the HTTP endpoint requires authorization.
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
type AwsIamConfig struct {
	// The signing region for AWS IAM authorization.
	// Experimental.
	SigningRegion *string `field:"required" json:"signingRegion" yaml:"signingRegion"`
	// The signing service name for AWS IAM authorization.
	// Experimental.
	SigningServiceName *string `field:"required" json:"signingServiceName" yaml:"signingServiceName"`
}

