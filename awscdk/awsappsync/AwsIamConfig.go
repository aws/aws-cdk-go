package awsappsync


// The authorization config in case the HTTP endpoint requires authorization.
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
type AwsIamConfig struct {
	// The signing region for AWS IAM authorization.
	SigningRegion *string `field:"required" json:"signingRegion" yaml:"signingRegion"`
	// The signing service name for AWS IAM authorization.
	SigningServiceName *string `field:"required" json:"signingServiceName" yaml:"signingServiceName"`
}

