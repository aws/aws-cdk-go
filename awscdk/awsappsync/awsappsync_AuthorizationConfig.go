package awsappsync


// Configuration of the API authorization modes.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//   var authFunction function
//
//
//   appsync.NewGraphqlApi(this, jsii.String("api"), &graphqlApiProps{
//   	name: jsii.String("api"),
//   	schema: appsync.schemaFile.fromAsset(path.join(__dirname, jsii.String("appsync.test.graphql"))),
//   	authorizationConfig: &authorizationConfig{
//   		defaultAuthorization: &authorizationMode{
//   			authorizationType: appsync.authorizationType_LAMBDA,
//   			lambdaAuthorizerConfig: &lambdaAuthorizerConfig{
//   				handler: authFunction,
//   			},
//   		},
//   	},
//   })
//
type AuthorizationConfig struct {
	// Additional authorization modes.
	AdditionalAuthorizationModes *[]*AuthorizationMode `field:"optional" json:"additionalAuthorizationModes" yaml:"additionalAuthorizationModes"`
	// Optional authorization configuration.
	DefaultAuthorization *AuthorizationMode `field:"optional" json:"defaultAuthorization" yaml:"defaultAuthorization"`
}

